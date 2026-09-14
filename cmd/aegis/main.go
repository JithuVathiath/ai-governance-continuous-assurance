package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/drift"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/iojson"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/ledger"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/policy"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/report"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/sarif"
)

const version = "0.1.0"

func main() {
	if err := run(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "aegis:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string) error {
	if len(args) == 0 {
		usage()
		return errors.New("a command is required")
	}
	switch args[0] {
	case "assess":
		return runAssess(ctx, args[1:])
	case "diff":
		return runDiff(args[1:])
	case "verify-ledger":
		return runVerifyLedger(ctx, args[1:])
	case "version":
		fmt.Println("aegis", version)
		return nil
	case "help", "-h", "--help":
		usage()
		return nil
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func runAssess(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet("assess", flag.ContinueOnError)
	systemPath := flags.String("system", "", "path to the AI system JSON")
	evidencePath := flags.String("evidence", "", "path to the evidence bundle JSON")
	policyDirectory := flags.String("policies", "policies", "directory containing policy packs")
	outputPath := flags.String("out", "assessment.json", "assessment JSON output")
	reportPath := flags.String("report", "assessment.md", "Markdown report output")
	sarifPath := flags.String("sarif", "assessment.sarif", "SARIF output")
	ledgerPath := flags.String("ledger", "aegis-ledger.db", "SQLite audit ledger")
	atValue := flags.String("at", "", "assessment time in RFC3339; defaults to current UTC time")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *systemPath == "" || *evidencePath == "" {
		return errors.New("--system and --evidence are required")
	}

	at := time.Now().UTC()
	if *atValue != "" {
		parsed, err := time.Parse(time.RFC3339, *atValue)
		if err != nil {
			return fmt.Errorf("parse --at: %w", err)
		}
		at = parsed
	}
	system, err := iojson.Load[domain.System](*systemPath)
	if err != nil {
		return err
	}
	bundle, err := iojson.Load[domain.EvidenceBundle](*evidencePath)
	if err != nil {
		return err
	}
	if bundle.SystemID != system.ID || bundle.Version != system.Version {
		return errors.New("evidence bundle identity does not match system identity")
	}
	evidence, err := policy.ResolveEvidence(*evidencePath, bundle)
	if err != nil {
		return err
	}
	packs, err := policy.LoadPacks(*policyDirectory)
	if err != nil {
		return err
	}
	assessment, err := policy.Assess(system, evidence, packs, at)
	if err != nil {
		return err
	}
	receipt, err := ledger.Append(ctx, *ledgerPath, assessment)
	if err != nil {
		return err
	}
	assessment.Ledger = &receipt
	if err := iojson.Write(*outputPath, assessment); err != nil {
		return err
	}
	if err := report.Write(*reportPath, assessment); err != nil {
		return err
	}
	if err := sarif.Write(*sarifPath, *systemPath, assessment); err != nil {
		return err
	}
	fmt.Printf("%s: %s — %.1f%% evidence coverage, %d finding(s)\n",
		assessment.System.ID, strings.ToUpper(assessment.Summary.Decision),
		assessment.Summary.EvidenceCoverage, len(assessment.Findings))
	return nil
}

func runDiff(args []string) error {
	flags := flag.NewFlagSet("diff", flag.ContinueOnError)
	beforePath := flags.String("before", "", "baseline assessment JSON")
	afterPath := flags.String("after", "", "changed assessment JSON")
	outputPath := flags.String("out", "drift.json", "drift JSON output")
	reportPath := flags.String("report", "drift.md", "drift Markdown report")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *beforePath == "" || *afterPath == "" {
		return errors.New("--before and --after are required")
	}
	before, err := iojson.Load[domain.Assessment](*beforePath)
	if err != nil {
		return err
	}
	after, err := iojson.Load[domain.Assessment](*afterPath)
	if err != nil {
		return err
	}
	driftReport, err := drift.Compare(before, after)
	if err != nil {
		return err
	}
	if err := iojson.Write(*outputPath, driftReport); err != nil {
		return err
	}
	if err := drift.WriteMarkdown(*reportPath, driftReport); err != nil {
		return err
	}
	fmt.Printf("drift: %d new, %d resolved, %d regressed, %d changed\n",
		driftReport.NewControls, driftReport.ResolvedControls,
		driftReport.RegressedControls, driftReport.ChangedControls)
	return nil
}

func runVerifyLedger(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet("verify-ledger", flag.ContinueOnError)
	ledgerPath := flags.String("ledger", "aegis-ledger.db", "SQLite audit ledger")
	outputPath := flags.String("out", "", "optional verification JSON output")
	if err := flags.Parse(args); err != nil {
		return err
	}
	verification, err := ledger.Verify(ctx, *ledgerPath)
	if err != nil {
		return err
	}
	if !verification.Valid {
		return fmt.Errorf("ledger verification failed at event %d: %s", verification.BrokenAt, verification.Explanation)
	}
	if *outputPath != "" {
		if err := iojson.Write(*outputPath, verification); err != nil {
			return err
		}
	}
	fmt.Printf("ledger valid: %d event(s), head %s\n", verification.Events, verification.HeadHash)
	return nil
}

func usage() {
	fmt.Print(`Aegis — evidence-led continuous assurance for AI systems

Usage:
  aegis assess --system SYSTEM.json --evidence EVIDENCE.json [options]
  aegis diff --before ASSESSMENT.json --after ASSESSMENT.json [options]
  aegis verify-ledger --ledger aegis-ledger.db [--out verification.json]
  aegis version

The output reports evidence readiness, not legal compliance.
`)
}

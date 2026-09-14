package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestRunEndToEnd(t *testing.T) {
	ctx := context.Background()
	root := filepath.Join("..", "..")
	directory := t.TempDir()
	ledgerPath := filepath.Join(directory, "ledger.db")
	v1 := filepath.Join(directory, "v1.json")
	v2 := filepath.Join(directory, "v2.json")

	assess := func(caseName, output string) {
		t.Helper()
		err := run(ctx, []string{
			"assess",
			"--system", filepath.Join(root, "cases", caseName, "system.json"),
			"--evidence", filepath.Join(root, "cases", caseName, "evidence.json"),
			"--policies", filepath.Join(root, "policies"),
			"--out", output,
			"--report", output + ".md",
			"--sarif", output + ".sarif",
			"--ledger", ledgerPath,
			"--at", "2026-09-13T12:00:00Z",
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	assess("recruitment-screening-v1", v1)
	assess("recruitment-screening-v2", v2)

	if err := run(ctx, []string{"diff", "--before", v1, "--after", v2, "--out", filepath.Join(directory, "drift.json"), "--report", filepath.Join(directory, "drift.md")}); err != nil {
		t.Fatal(err)
	}
	verificationPath := filepath.Join(directory, "verification.json")
	if err := run(ctx, []string{"verify-ledger", "--ledger", ledgerPath, "--out", verificationPath}); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{v1, v2, v1 + ".md", v2 + ".sarif", verificationPath} {
		if info, err := os.Stat(path); err != nil || info.Size() == 0 {
			t.Fatalf("expected non-empty output %s: %v", path, err)
		}
	}
}

func TestRunCommandsAndErrors(t *testing.T) {
	ctx := context.Background()
	for _, args := range [][]string{{"help"}, {"version"}} {
		if err := run(ctx, args); err != nil {
			t.Fatalf("%v failed: %v", args, err)
		}
	}
	for _, args := range [][]string{{}, {"unknown"}, {"assess"}, {"diff"}, {"assess", "--system", "x", "--evidence", "y", "--at", "invalid"}} {
		if err := run(ctx, args); err == nil {
			t.Fatalf("%v should fail", args)
		}
	}
}

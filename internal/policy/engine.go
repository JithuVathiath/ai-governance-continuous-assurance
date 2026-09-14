package policy

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/url"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/classify"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/iojson"
)

const disclaimer = "Aegis reports evidence readiness and traceability. It is not a legal opinion, certification, or determination of regulatory compliance."

// LoadPacks loads and fingerprints every JSON policy pack in a directory.
func LoadPacks(directory string) ([]domain.PolicyPack, error) {
	paths, err := filepath.Glob(filepath.Join(directory, "*.json"))
	if err != nil {
		return nil, fmt.Errorf("list policy packs: %w", err)
	}
	if len(paths) == 0 {
		return nil, fmt.Errorf("no policy packs found in %s", directory)
	}
	sort.Strings(paths)

	packs := make([]domain.PolicyPack, 0, len(paths))
	for _, path := range paths {
		pack, loadErr := iojson.Load[domain.PolicyPack](path)
		if loadErr != nil {
			return nil, loadErr
		}
		digest, digestErr := iojson.Digest(path)
		if digestErr != nil {
			return nil, digestErr
		}
		pack.Digest = digest
		if err := validatePack(pack); err != nil {
			return nil, fmt.Errorf("validate %s: %w", path, err)
		}
		packs = append(packs, pack)
	}
	return packs, nil
}

// ResolveEvidence verifies referenced artefacts and annotates their integrity.
func ResolveEvidence(bundlePath string, bundle domain.EvidenceBundle) ([]domain.Evidence, error) {
	base, err := filepath.Abs(filepath.Dir(bundlePath))
	if err != nil {
		return nil, fmt.Errorf("resolve evidence directory: %w", err)
	}
	resolvedBase, err := filepath.EvalSymlinks(base)
	if err != nil {
		return nil, fmt.Errorf("resolve evidence directory: %w", err)
	}
	items := slices.Clone(bundle.Items)
	seen := map[string]bool{}
	for index := range items {
		item := &items[index]
		if item.ID == "" || seen[item.ID] {
			return nil, fmt.Errorf("evidence IDs must be present and unique: %q", item.ID)
		}
		seen[item.ID] = true
		if item.Type == "" || item.Title == "" || item.Owner == "" || item.Artifact == "" {
			return nil, fmt.Errorf("evidence %s requires type, title, owner, and artifact", item.ID)
		}
		if !slices.Contains([]string{"approved", "draft", "in-review", "rejected"}, item.Status) {
			return nil, fmt.Errorf("evidence %s has unsupported status %q", item.ID, item.Status)
		}
		if _, parseErr := time.Parse("2006-01-02", item.CollectedAt); parseErr != nil {
			return nil, fmt.Errorf("evidence %s has invalid collected_at: %w", item.ID, parseErr)
		}
		if item.ValidUntil != "" {
			if _, parseErr := time.Parse("2006-01-02", item.ValidUntil); parseErr != nil {
				return nil, fmt.Errorf("evidence %s has invalid valid_until: %w", item.ID, parseErr)
			}
		}
		path := filepath.Clean(filepath.Join(resolvedBase, item.Artifact))
		resolvedPath, resolveErr := filepath.EvalSymlinks(path)
		if resolveErr != nil {
			return nil, fmt.Errorf("resolve evidence %s: %w", item.ID, resolveErr)
		}
		relative, relativeErr := filepath.Rel(resolvedBase, resolvedPath)
		if relativeErr != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
			return nil, fmt.Errorf("evidence %s escapes its bundle directory", item.ID)
		}
		digest, digestErr := iojson.Digest(resolvedPath)
		if digestErr != nil {
			return nil, fmt.Errorf("verify evidence %s: %w", item.ID, digestErr)
		}
		if item.SHA256 != "" && item.SHA256 != "auto" && item.SHA256 != digest {
			return nil, fmt.Errorf("evidence %s digest mismatch", item.ID)
		}
		item.SHA256 = digest
		item.Integrity = "verified"
	}
	return items, nil
}

// Assess evaluates a system and evidence bundle against versioned policy packs.
func Assess(system domain.System, evidence []domain.Evidence, packs []domain.PolicyPack, at time.Time) (domain.Assessment, error) {
	if err := validateSystem(system); err != nil {
		return domain.Assessment{}, err
	}
	classification := classify.Evaluate(system)

	assessment := domain.Assessment{
		SchemaVersion:  "1.0.0",
		GeneratedAt:    at.UTC().Format(time.RFC3339),
		System:         system,
		Classification: classification,
		Evidence:       evidence,
		Disclaimer:     disclaimer,
		Graph: domain.EvidenceGraph{
			Nodes: []domain.GraphNode{{ID: "system:" + system.ID, Type: "system", Label: system.Name}},
			Edges: []domain.GraphEdge{},
		},
	}

	for _, pack := range packs {
		assessment.PolicyVersions = append(assessment.PolicyVersions, domain.PolicyVersion{
			PackID: pack.ID, Framework: pack.Framework, Version: pack.Version,
			AsOf: pack.AsOf, SourceURL: pack.SourceURL, PackSHA256: pack.Digest,
		})
		packNode := "framework:" + pack.ID
		assessment.Graph.Nodes = append(assessment.Graph.Nodes, domain.GraphNode{
			ID: packNode, Type: "framework", Label: pack.Framework + " " + pack.Version,
		})
		assessment.Graph.Edges = append(assessment.Graph.Edges, domain.GraphEdge{
			From: "system:" + system.ID, To: packNode, Relation: "assessed_against",
		})

		for _, rule := range pack.Rules {
			result := evaluateRule(system, classification.Tier, evidence, pack, rule, at)
			assessment.Controls = append(assessment.Controls, result)
			if !result.Applicable {
				continue
			}
			controlNode := "control:" + pack.ID + ":" + rule.ID
			assessment.Graph.Nodes = append(assessment.Graph.Nodes, domain.GraphNode{
				ID: controlNode, Type: "control", Label: rule.Title,
				Meta: map[string]string{"status": result.Status, "severity": result.Severity},
			})
			assessment.Graph.Edges = append(assessment.Graph.Edges, domain.GraphEdge{
				From: packNode, To: controlNode, Relation: "defines",
			})

			for _, requirement := range result.Requirements {
				for _, evidenceID := range requirement.EvidenceIDs {
					evidenceNode := "evidence:" + evidenceID
					ensureEvidenceNode(&assessment.Graph, evidenceNode, evidenceID, evidence)
					assessment.Graph.Edges = append(assessment.Graph.Edges, domain.GraphEdge{
						From: controlNode, To: evidenceNode, Relation: "supported_by",
					})
				}
			}

			if result.Status != "supported" {
				finding := findingFor(result)
				assessment.Findings = append(assessment.Findings, finding)
				findingNode := "finding:" + finding.ID
				assessment.Graph.Nodes = append(assessment.Graph.Nodes, domain.GraphNode{
					ID: findingNode, Type: "finding", Label: finding.Title,
					Meta: map[string]string{"severity": finding.Severity, "status": finding.Status},
				})
				assessment.Graph.Edges = append(assessment.Graph.Edges, domain.GraphEdge{
					From: controlNode, To: findingNode, Relation: "raises",
				})
			}
		}
	}

	if classification.DeclarationMismatch {
		assessment.Findings = append(assessment.Findings, domain.Finding{
			ID: "risk-tier-declaration-mismatch", ControlID: "classification",
			Framework: "Aegis", Severity: "high", Status: "mismatch",
			Title: "Declared risk tier differs from engine routing",
			Message: fmt.Sprintf("Declared %q; engine routed %q. Independent review is required.",
				classification.DeclaredTier, classification.Tier),
		})
	}
	assessment.Summary = summarise(assessment.Controls, assessment.Findings)
	assessment.AssessmentID = assessmentIdentity(assessment)
	return assessment, nil
}

func evaluateRule(system domain.System, tier string, evidence []domain.Evidence, pack domain.PolicyPack, rule domain.Rule, at time.Time) domain.ControlResult {
	result := domain.ControlResult{
		ID: rule.ID, Framework: pack.Framework, PackID: pack.ID, Title: rule.Title,
		SourceRef: rule.SourceRef, Severity: rule.Severity,
	}
	if !applies(rule.ApplicableWhen, system, tier) {
		result.Status = "not-applicable"
		result.Reason = "Applicability conditions did not match the assessed system."
		return result
	}
	result.Applicable = true

	statuses := make([]string, 0, len(rule.Requirements))
	for _, requirement := range rule.Requirements {
		requirementResult := evaluateRequirement(requirement, evidence, at)
		result.Requirements = append(result.Requirements, requirementResult)
		statuses = append(statuses, requirementResult.Status)
	}
	result.Status = aggregateStatus(statuses)
	result.Reason = statusReason(result.Status)
	return result
}

func evaluateRequirement(requirement domain.Requirement, evidence []domain.Evidence, at time.Time) domain.RequirementResult {
	result := domain.RequirementResult{ID: requirement.ID, Title: requirement.Title}
	minimum := max(requirement.Minimum, 1)
	approved := 0
	draft := 0
	expired := 0

	for _, item := range evidence {
		if !slices.Contains(requirement.EvidenceTypes, item.Type) {
			continue
		}
		result.EvidenceIDs = append(result.EvidenceIDs, item.ID)
		if item.ValidUntil != "" {
			validUntil, err := time.Parse("2006-01-02", item.ValidUntil)
			if err == nil && at.After(validUntil.Add(24*time.Hour-time.Nanosecond)) {
				expired++
				continue
			}
		}
		switch item.Status {
		case "approved":
			if item.Integrity == "verified" {
				approved++
			}
		case "draft", "in-review":
			draft++
		}
	}

	sort.Strings(result.EvidenceIDs)
	switch {
	case approved >= minimum:
		result.Status = "supported"
		result.Reason = fmt.Sprintf("%d verified, approved evidence item(s) satisfy the minimum of %d.", approved, minimum)
	case expired > 0 && approved == 0 && draft == 0:
		result.Status = "expired"
		result.Reason = fmt.Sprintf("%d matching evidence item(s) are outside their validity period.", expired)
	case approved+draft > 0:
		result.Status = "partial"
		result.Reason = fmt.Sprintf("Evidence exists but only %d of %d required item(s) are verified and approved.", approved, minimum)
	default:
		result.Status = "missing"
		result.Reason = "No matching evidence was supplied."
	}
	return result
}

func applies(conditions domain.Conditions, system domain.System, tier string) bool {
	return matchScalar(conditions.RiskTiers, tier) &&
		matchScalar(conditions.Domains, system.Domain) &&
		matchScalar(conditions.Roles, system.Role) &&
		matchAny(conditions.Geographies, system.Geographies) &&
		matchAny(conditions.Capabilities, system.Capabilities) &&
		matchAny(conditions.RiskFactors, system.RiskFactors) &&
		matchScalar(conditions.LifecycleStage, system.LifecycleStage)
}

func matchScalar(expected []string, actual string) bool {
	if len(expected) == 0 {
		return true
	}
	for _, value := range expected {
		if strings.EqualFold(value, actual) {
			return true
		}
	}
	return false
}

func matchAny(expected, actual []string) bool {
	if len(expected) == 0 {
		return true
	}
	for _, expectedValue := range expected {
		for _, actualValue := range actual {
			if strings.EqualFold(expectedValue, actualValue) {
				return true
			}
		}
	}
	return false
}

func aggregateStatus(statuses []string) string {
	if len(statuses) == 0 {
		return "supported"
	}
	for _, status := range []string{"missing", "expired", "partial"} {
		if slices.Contains(statuses, status) {
			return status
		}
	}
	return "supported"
}

func statusReason(status string) string {
	switch status {
	case "supported":
		return "Every requirement has sufficient verified, approved evidence."
	case "partial":
		return "At least one requirement has draft or insufficient approved evidence."
	case "expired":
		return "At least one requirement relies only on expired evidence."
	default:
		return "At least one requirement has no matching evidence."
	}
}

func findingFor(result domain.ControlResult) domain.Finding {
	return domain.Finding{
		ID:        strings.ToLower(result.PackID + "-" + result.ID + "-" + result.Status),
		ControlID: result.ID,
		Framework: result.Framework,
		Severity:  result.Severity,
		Status:    result.Status,
		Title:     result.Title + ": " + result.Status,
		Message:   result.Reason,
	}
}

func summarise(controls []domain.ControlResult, findings []domain.Finding) domain.Summary {
	summary := domain.Summary{}
	criticalGap := false
	for _, control := range controls {
		if !control.Applicable {
			continue
		}
		summary.Applicable++
		switch control.Status {
		case "supported":
			summary.Supported++
		case "partial":
			summary.Partial++
		case "expired":
			summary.Expired++
		default:
			summary.Missing++
		}
		if control.Status != "supported" && (control.Severity == "critical" || control.Severity == "high") {
			criticalGap = true
		}
	}
	if summary.Applicable > 0 {
		summary.EvidenceCoverage = math.Round(float64(summary.Supported)/float64(summary.Applicable)*1000) / 10
	}
	switch {
	case criticalGap || slices.ContainsFunc(findings, func(f domain.Finding) bool { return f.ID == "risk-tier-declaration-mismatch" }):
		summary.Decision = "hold"
	case summary.Partial+summary.Missing+summary.Expired > 0:
		summary.Decision = "conditional-review"
	default:
		summary.Decision = "ready-for-independent-review"
	}
	return summary
}

func ensureEvidenceNode(graph *domain.EvidenceGraph, nodeID, evidenceID string, evidence []domain.Evidence) {
	if slices.ContainsFunc(graph.Nodes, func(node domain.GraphNode) bool { return node.ID == nodeID }) {
		return
	}
	label := evidenceID
	meta := map[string]string{}
	for _, item := range evidence {
		if item.ID == evidenceID {
			label = item.Title
			meta = map[string]string{"sha256": item.SHA256, "status": item.Status, "integrity": item.Integrity}
			break
		}
	}
	graph.Nodes = append(graph.Nodes, domain.GraphNode{ID: nodeID, Type: "evidence", Label: label, Meta: meta})
}

func validatePack(pack domain.PolicyPack) error {
	if pack.ID == "" || pack.Framework == "" || pack.Version == "" || pack.SourceURL == "" {
		return errors.New("pack identity, version, framework, and source URL are required")
	}
	if _, err := time.Parse("2006-01-02", pack.AsOf); err != nil {
		return fmt.Errorf("pack as_of must be an ISO date: %w", err)
	}
	source, err := url.ParseRequestURI(pack.SourceURL)
	if err != nil || source.Scheme != "https" || source.Host == "" {
		return errors.New("pack source URL must be an absolute HTTPS URL")
	}
	seen := map[string]bool{}
	for _, rule := range pack.Rules {
		if rule.ID == "" || seen[rule.ID] {
			return fmt.Errorf("rule IDs must be present and unique: %q", rule.ID)
		}
		seen[rule.ID] = true
		if rule.Title == "" || !slices.Contains([]string{"critical", "high", "medium", "low"}, rule.Severity) || len(rule.Requirements) == 0 {
			return fmt.Errorf("rule %s is incomplete", rule.ID)
		}
		requirements := map[string]bool{}
		for _, requirement := range rule.Requirements {
			if requirement.ID == "" || requirements[requirement.ID] || requirement.Title == "" || len(requirement.EvidenceTypes) == 0 || requirement.Minimum < 1 {
				return fmt.Errorf("rule %s has an invalid or duplicate requirement %q", rule.ID, requirement.ID)
			}
			requirements[requirement.ID] = true
		}
	}
	return nil
}

func validateSystem(system domain.System) error {
	if system.ID == "" || system.Name == "" || system.Version == "" || system.Owner == "" {
		return errors.New("system id, name, version, and owner are required")
	}
	if system.Domain == "" || system.Role == "" || system.LifecycleStage == "" {
		return errors.New("system domain, role, and lifecycle stage are required")
	}
	return nil
}

func assessmentIdentity(assessment domain.Assessment) string {
	assessment.AssessmentID = ""
	assessment.Ledger = nil
	payload, _ := json.Marshal(assessment)
	sum := sha256.Sum256(payload)
	return "asg-" + hex.EncodeToString(sum[:8])
}

// MarshalCanonical produces the stable assessment payload used by the ledger.
func MarshalCanonical(assessment domain.Assessment) ([]byte, error) {
	assessment.Ledger = nil
	return json.Marshal(assessment)
}

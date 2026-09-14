package policy

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

func TestAssessProducesTraceableHold(t *testing.T) {
	system := validSystem()
	evidence := []domain.Evidence{
		{ID: "approved", Type: "risk_assessment", Status: "approved", Integrity: "verified"},
		{ID: "expired", Type: "security_assessment", Status: "approved", Integrity: "verified", ValidUntil: "2025-01-01"},
	}
	pack := domain.PolicyPack{ID: "pack", Framework: "Framework", Version: "1", AsOf: "2026-01-01", SourceURL: "https://example.test", Digest: "abc", Rules: []domain.Rule{
		{ID: "R1", Title: "Risk", SourceRef: "one", Severity: "high", ApplicableWhen: domain.Conditions{RiskTiers: []string{"high"}}, Requirements: []domain.Requirement{{ID: "risk", Title: "Risk", EvidenceTypes: []string{"risk_assessment"}, Minimum: 1}}},
		{ID: "R2", Title: "Security", SourceRef: "two", Severity: "high", ApplicableWhen: domain.Conditions{RiskTiers: []string{"high"}}, Requirements: []domain.Requirement{{ID: "security", Title: "Security", EvidenceTypes: []string{"security_assessment"}, Minimum: 1}}},
	}}
	at := time.Date(2026, 9, 13, 12, 0, 0, 0, time.UTC)
	assessment, err := Assess(system, evidence, []domain.PolicyPack{pack}, at)
	if err != nil {
		t.Fatal(err)
	}
	if assessment.Summary.Decision != "hold" || assessment.Summary.Supported != 1 || assessment.Summary.Expired != 1 {
		t.Fatalf("unexpected summary: %#v", assessment.Summary)
	}
	if len(assessment.Graph.Nodes) < 5 || len(assessment.Findings) != 1 {
		t.Fatalf("expected evidence lineage and one finding, got %#v", assessment)
	}

	again, err := Assess(system, evidence, []domain.PolicyPack{pack}, at)
	if err != nil || assessment.AssessmentID != again.AssessmentID {
		t.Fatal("assessment identity must be deterministic")
	}
}

func TestResolveEvidenceVerifiesDigestAndRejectsEscape(t *testing.T) {
	directory := t.TempDir()
	artifact := filepath.Join(directory, "record.md")
	if err := os.WriteFile(artifact, []byte("evidence"), 0o600); err != nil {
		t.Fatal(err)
	}
	bundlePath := filepath.Join(directory, "evidence.json")
	valid := domain.Evidence{ID: "one", Type: "record", Title: "Record", Status: "approved", Owner: "Owner", Artifact: "record.md", SHA256: "auto", CollectedAt: "2026-01-01"}
	items, err := ResolveEvidence(bundlePath, domain.EvidenceBundle{Items: []domain.Evidence{valid}})
	if err != nil || items[0].Integrity != "verified" || len(items[0].SHA256) != 64 {
		t.Fatalf("unexpected resolution: %#v, %v", items, err)
	}

	escape := valid
	escape.ID = "escape"
	escape.Artifact = "../outside.md"
	_, err = ResolveEvidence(bundlePath, domain.EvidenceBundle{Items: []domain.Evidence{escape}})
	if err == nil {
		t.Fatal("expected path escape to fail")
	}
	invalidDigest := valid
	invalidDigest.SHA256 = "wrong"
	_, err = ResolveEvidence(bundlePath, domain.EvidenceBundle{Items: []domain.Evidence{invalidDigest}})
	if err == nil {
		t.Fatal("expected digest mismatch to fail")
	}
}

func TestResolveEvidenceRejectsInvalidMetadata(t *testing.T) {
	directory := t.TempDir()
	if err := os.WriteFile(filepath.Join(directory, "record.md"), []byte("evidence"), 0o600); err != nil {
		t.Fatal(err)
	}
	bundlePath := filepath.Join(directory, "evidence.json")
	item := domain.Evidence{ID: "one", Type: "record", Title: "Record", Status: "invented", Owner: "Owner", Artifact: "record.md", CollectedAt: "2026-01-01"}
	if _, err := ResolveEvidence(bundlePath, domain.EvidenceBundle{Items: []domain.Evidence{item}}); err == nil {
		t.Fatal("expected invalid status")
	}
	item.Status = "approved"
	item.CollectedAt = "yesterday"
	if _, err := ResolveEvidence(bundlePath, domain.EvidenceBundle{Items: []domain.Evidence{item}}); err == nil {
		t.Fatal("expected invalid date")
	}
}

func TestResolveEvidenceRejectsSymlinkEscape(t *testing.T) {
	directory := t.TempDir()
	outside := filepath.Join(t.TempDir(), "outside.md")
	if err := os.WriteFile(outside, []byte("outside"), 0o600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(directory, "linked.md")
	if err := os.Symlink(outside, link); err != nil {
		t.Skipf("symbolic links are unavailable: %v", err)
	}
	item := domain.Evidence{ID: "one", Type: "record", Title: "Record", Status: "approved", Owner: "Owner", Artifact: "linked.md", CollectedAt: "2026-01-01"}
	if _, err := ResolveEvidence(filepath.Join(directory, "evidence.json"), domain.EvidenceBundle{Items: []domain.Evidence{item}}); err == nil {
		t.Fatal("expected symbolic-link escape to fail")
	}
}

func TestAssessRejectsInvalidSystem(t *testing.T) {
	if _, err := Assess(domain.System{}, nil, nil, time.Now()); err == nil {
		t.Fatal("expected invalid system to fail")
	}
}

func TestLoadPacksAndPackValidation(t *testing.T) {
	packs, err := LoadPacks(filepath.Join("..", "..", "policies"))
	if err != nil {
		t.Fatal(err)
	}
	if len(packs) != 3 {
		t.Fatalf("got %d packs, want 3", len(packs))
	}
	for _, pack := range packs {
		if len(pack.Digest) != 64 {
			t.Fatalf("pack %s was not fingerprinted", pack.ID)
		}
	}
	if _, err := LoadPacks(t.TempDir()); err == nil {
		t.Fatal("expected empty directory to fail")
	}

	invalid := domain.PolicyPack{ID: "p", Framework: "F", Version: "1", AsOf: "not-a-date", SourceURL: "http://example.test"}
	if err := validatePack(invalid); err == nil {
		t.Fatal("expected invalid metadata")
	}
	invalid.AsOf = "2026-01-01"
	invalid.SourceURL = "https://example.test"
	invalid.Rules = []domain.Rule{{ID: "R", Title: "Rule", Severity: "unknown", Requirements: []domain.Requirement{{ID: "Q", Title: "Requirement", EvidenceTypes: []string{"record"}, Minimum: 1}}}}
	if err := validatePack(invalid); err == nil {
		t.Fatal("expected invalid severity")
	}
	invalid.Rules[0].Severity = "high"
	invalid.Rules[0].Requirements[0].Minimum = 0
	if err := validatePack(invalid); err == nil {
		t.Fatal("expected invalid requirement")
	}
}

func TestEvidenceStatusesAndDecisionGates(t *testing.T) {
	at := time.Date(2026, 9, 13, 12, 0, 0, 0, time.UTC)
	requirement := domain.Requirement{ID: "one", Title: "One", EvidenceTypes: []string{"record"}, Minimum: 1}
	tests := []struct {
		name     string
		evidence []domain.Evidence
		want     string
	}{
		{"supported", []domain.Evidence{{Type: "record", Status: "approved", Integrity: "verified"}}, "supported"},
		{"partial", []domain.Evidence{{Type: "record", Status: "draft", Integrity: "verified"}}, "partial"},
		{"expired", []domain.Evidence{{Type: "record", Status: "approved", Integrity: "verified", ValidUntil: "2025-01-01"}}, "expired"},
		{"missing", nil, "missing"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := evaluateRequirement(requirement, test.evidence, at).Status; got != test.want {
				t.Fatalf("got %s, want %s", got, test.want)
			}
		})
	}
	ready := summarise([]domain.ControlResult{{Applicable: true, Status: "supported", Severity: "high"}}, nil)
	conditional := summarise([]domain.ControlResult{{Applicable: true, Status: "missing", Severity: "low"}}, nil)
	if ready.Decision != "ready-for-independent-review" || conditional.Decision != "conditional-review" {
		t.Fatalf("unexpected decisions: %#v %#v", ready, conditional)
	}
	if matchAny([]string{"EU"}, []string{"US"}) || !matchAny(nil, nil) || !matchScalar(nil, "anything") {
		t.Fatal("applicability matching is incorrect")
	}
}

func TestMarshalCanonicalExcludesLedger(t *testing.T) {
	receipt := &domain.LedgerReceipt{Sequence: 1}
	data, err := MarshalCanonical(domain.Assessment{AssessmentID: "one", Ledger: receipt})
	if err != nil {
		t.Fatal(err)
	}
	if string(data) == "" || contains(string(data), "\"ledger\"") {
		t.Fatalf("canonical payload includes ledger: %s", data)
	}
}

func validSystem() domain.System {
	return domain.System{ID: "system", Name: "System", Version: "1", Owner: "Owner", Role: "provider", Domain: "employment", LifecycleStage: "pilot"}
}

func contains(value, fragment string) bool {
	for i := 0; i+len(fragment) <= len(value); i++ {
		if value[i:i+len(fragment)] == fragment {
			return true
		}
	}
	return false
}

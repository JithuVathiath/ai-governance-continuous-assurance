package sarif

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

func TestWriteProducesSARIF(t *testing.T) {
	path := filepath.Join(t.TempDir(), "result.sarif")
	assessment := domain.Assessment{Findings: []domain.Finding{{ID: "gap", ControlID: "C1", Framework: "Framework", Severity: "high", Title: "Missing control", Message: "Evidence is missing."}}}
	if err := Write(path, "system.json", assessment); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var document map[string]any
	if err := json.Unmarshal(data, &document); err != nil {
		t.Fatal(err)
	}
	if document["version"] != "2.1.0" {
		t.Fatalf("unexpected SARIF: %#v", document)
	}
}

func TestSARIFLevels(t *testing.T) {
	tests := map[string]string{"critical": "error", "high": "error", "medium": "warning", "low": "note"}
	for severity, want := range tests {
		if got := sarifLevel(severity); got != want {
			t.Fatalf("severity %s produced %s, want %s", severity, got, want)
		}
	}
}

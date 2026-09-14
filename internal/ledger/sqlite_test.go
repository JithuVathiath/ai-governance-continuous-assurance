package ledger

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

func TestAppendAndVerifyDetectsTampering(t *testing.T) {
	ctx := context.Background()
	path := filepath.Join(t.TempDir(), "ledger.db")
	assessment := domain.Assessment{AssessmentID: "asg-one", GeneratedAt: "2026-09-13T12:00:00Z", System: domain.System{ID: "one", Version: "1"}}
	first, err := Append(ctx, path, assessment)
	if err != nil {
		t.Fatal(err)
	}
	assessment.AssessmentID = "asg-two"
	assessment.GeneratedAt = "2026-09-13T12:01:00Z"
	second, err := Append(ctx, path, assessment)
	if err != nil {
		t.Fatal(err)
	}
	if first.Sequence != 1 || second.Sequence != 2 || second.PreviousHash != first.EventHash {
		t.Fatalf("invalid receipts: %#v %#v", first, second)
	}
	if _, err := Append(ctx, path, assessment); err == nil {
		t.Fatal("expected duplicate assessment to be rejected")
	}
	verification, err := Verify(ctx, path)
	if err != nil || !verification.Valid || verification.Events != 2 {
		t.Fatalf("unexpected verification: %#v, %v", verification, err)
	}

	database, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := database.ExecContext(ctx, "UPDATE events SET payload_json = ? WHERE sequence = 1", "tampered"); err != nil {
		t.Fatal(err)
	}
	if err := database.Close(); err != nil {
		t.Fatal(err)
	}
	verification, err = Verify(ctx, path)
	if err != nil || verification.Valid || verification.BrokenAt != 1 {
		t.Fatalf("tampering was not detected: %#v, %v", verification, err)
	}
}

func TestVerifyRejectsMissingLedger(t *testing.T) {
	if _, err := Verify(context.Background(), filepath.Join(t.TempDir(), "missing.db")); err == nil {
		t.Fatal("expected a missing ledger to fail")
	}
}

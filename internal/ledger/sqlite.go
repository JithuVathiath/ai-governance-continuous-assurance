package ledger

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/policy"
	_ "modernc.org/sqlite"
)

const genesis = "GENESIS"

// Verification captures the result of a complete ledger-chain audit.
type Verification struct {
	Valid       bool   `json:"valid"`
	Events      int64  `json:"events"`
	HeadHash    string `json:"head_hash"`
	BrokenAt    int64  `json:"broken_at,omitempty"`
	Explanation string `json:"explanation"`
}

// Append records an assessment in an append-only SHA-256 hash chain.
func Append(ctx context.Context, path string, assessment domain.Assessment) (domain.LedgerReceipt, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return domain.LedgerReceipt{}, fmt.Errorf("create ledger directory: %w", err)
	}
	database, err := sql.Open("sqlite", path)
	if err != nil {
		return domain.LedgerReceipt{}, fmt.Errorf("open ledger: %w", err)
	}
	defer database.Close()
	if err := migrate(ctx, database); err != nil {
		return domain.LedgerReceipt{}, err
	}

	payload, err := policy.MarshalCanonical(assessment)
	if err != nil {
		return domain.LedgerReceipt{}, fmt.Errorf("encode assessment payload: %w", err)
	}
	payloadHash := digest(payload)

	transaction, err := database.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return domain.LedgerReceipt{}, fmt.Errorf("begin ledger transaction: %w", err)
	}
	defer transaction.Rollback()

	var sequence int64
	previousHash := genesis
	row := transaction.QueryRowContext(ctx, "SELECT sequence, event_hash FROM events ORDER BY sequence DESC LIMIT 1")
	if scanErr := row.Scan(&sequence, &previousHash); scanErr != nil && !errors.Is(scanErr, sql.ErrNoRows) {
		return domain.LedgerReceipt{}, fmt.Errorf("read ledger head: %w", scanErr)
	}
	sequence++
	eventType := "assessment.completed"
	eventHash := calculateHash(sequence, assessment.GeneratedAt, assessment.System.ID, eventType, payloadHash, previousHash)

	_, err = transaction.ExecContext(ctx, `
		INSERT INTO events(sequence, occurred_at, system_id, system_version, event_type,
			assessment_id, payload_json, payload_hash, previous_hash, event_hash)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		sequence, assessment.GeneratedAt, assessment.System.ID, assessment.System.Version, eventType,
		assessment.AssessmentID, string(payload), payloadHash, previousHash, eventHash,
	)
	if err != nil {
		return domain.LedgerReceipt{}, fmt.Errorf("append ledger event: %w", err)
	}
	if err := transaction.Commit(); err != nil {
		return domain.LedgerReceipt{}, fmt.Errorf("commit ledger event: %w", err)
	}

	return domain.LedgerReceipt{
		Sequence: sequence, EventHash: eventHash, PreviousHash: previousHash, PayloadHash: payloadHash,
	}, nil
}

// Verify recomputes every payload and chain hash from genesis to the current head.
func Verify(ctx context.Context, path string) (Verification, error) {
	if _, err := os.Stat(path); err != nil {
		return Verification{}, fmt.Errorf("open ledger: %w", err)
	}
	database, err := sql.Open("sqlite", path)
	if err != nil {
		return Verification{}, fmt.Errorf("open ledger: %w", err)
	}
	defer database.Close()
	rows, err := database.QueryContext(ctx, `
		SELECT sequence, occurred_at, system_id, event_type, payload_json,
			payload_hash, previous_hash, event_hash
		FROM events ORDER BY sequence`)
	if err != nil {
		return Verification{}, fmt.Errorf("read ledger: %w", err)
	}
	defer rows.Close()

	verification := Verification{Valid: true, HeadHash: genesis, Explanation: "Hash chain is valid."}
	expectedPrevious := genesis
	for rows.Next() {
		var sequence int64
		var occurredAt, systemID, eventType, payloadJSON, payloadHash, previousHash, eventHash string
		if err := rows.Scan(&sequence, &occurredAt, &systemID, &eventType, &payloadJSON, &payloadHash, &previousHash, &eventHash); err != nil {
			return Verification{}, fmt.Errorf("scan ledger event: %w", err)
		}
		verification.Events++
		calculatedPayload := digest([]byte(payloadJSON))
		calculatedEvent := calculateHash(sequence, occurredAt, systemID, eventType, payloadHash, previousHash)
		if payloadHash != calculatedPayload || previousHash != expectedPrevious || eventHash != calculatedEvent {
			verification.Valid = false
			verification.BrokenAt = sequence
			verification.Explanation = "Payload or chain hash mismatch detected."
			return verification, nil
		}
		expectedPrevious = eventHash
		verification.HeadHash = eventHash
	}
	if err := rows.Err(); err != nil {
		return Verification{}, fmt.Errorf("iterate ledger: %w", err)
	}
	if verification.Events == 0 {
		verification.Valid = false
		verification.Explanation = "Ledger contains no events."
	}
	return verification, nil
}

func migrate(ctx context.Context, database *sql.DB) error {
	_, err := database.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS events (
			sequence INTEGER PRIMARY KEY,
			occurred_at TEXT NOT NULL,
			system_id TEXT NOT NULL,
			system_version TEXT NOT NULL,
			event_type TEXT NOT NULL,
			assessment_id TEXT NOT NULL,
			payload_json TEXT NOT NULL,
			payload_hash TEXT NOT NULL,
			previous_hash TEXT NOT NULL,
			event_hash TEXT NOT NULL UNIQUE
		);
		CREATE INDEX IF NOT EXISTS idx_events_system ON events(system_id, sequence);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_events_assessment ON events(assessment_id);
	`)
	if err != nil {
		return fmt.Errorf("migrate ledger: %w", err)
	}
	return nil
}

func calculateHash(sequence int64, occurredAt, systemID, eventType, payloadHash, previousHash string) string {
	canonical := strings.Join([]string{
		fmt.Sprintf("%d", sequence), occurredAt, systemID, eventType, payloadHash, previousHash,
	}, "|")
	return digest([]byte(canonical))
}

func digest(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

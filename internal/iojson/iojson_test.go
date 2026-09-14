package iojson

import (
	"os"
	"path/filepath"
	"testing"
)

type sample struct {
	Name string `json:"name"`
}

func TestWriteLoadAndDigest(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", "sample.json")
	if err := Write(path, sample{Name: "Aegis"}); err != nil {
		t.Fatal(err)
	}
	value, err := Load[sample](path)
	if err != nil || value.Name != "Aegis" {
		t.Fatalf("unexpected load: %#v, %v", value, err)
	}
	digest, err := Digest(path)
	if err != nil || len(digest) != 64 {
		t.Fatalf("unexpected digest: %q, %v", digest, err)
	}
}

func TestLoadRejectsUnknownAndTrailingData(t *testing.T) {
	directory := t.TempDir()
	tests := map[string]string{
		"unknown":  `{"name":"Aegis","extra":true}`,
		"trailing": `{"name":"Aegis"}{"name":"other"}`,
	}
	for name, content := range tests {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(directory, name+".json")
			if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
				t.Fatal(err)
			}
			if _, err := Load[sample](path); err == nil {
				t.Fatal("expected strict decode failure")
			}
		})
	}
}

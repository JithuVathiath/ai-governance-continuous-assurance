package iojson

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Load decodes strict JSON and rejects unknown fields.
func Load[T any](path string) (T, error) {
	var value T
	file, err := openScoped(path)
	if err != nil {
		return value, fmt.Errorf("open %s: %w", path, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&value); err != nil {
		return value, fmt.Errorf("decode %s: %w", path, err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return value, fmt.Errorf("decode %s: trailing JSON content", path)
	}
	return value, nil
}

// Write atomically writes indented JSON.
func Write(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("encode %s: %w", path, err)
	}
	data = append(data, '\n')
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	temporary := path + ".tmp"
	if err := os.WriteFile(temporary, data, 0o600); err != nil {
		return fmt.Errorf("write temporary output: %w", err)
	}
	if err := os.Rename(temporary, path); err != nil {
		return fmt.Errorf("publish output: %w", err)
	}
	return nil
}

// Digest returns the SHA-256 digest of a file.
func Digest(path string) (string, error) {
	file, err := openScoped(path)
	if err != nil {
		return "", fmt.Errorf("open %s: %w", path, err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

func openScoped(path string) (io.ReadCloser, error) {
	absolute, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	root, err := os.OpenRoot(filepath.Dir(absolute))
	if err != nil {
		return nil, err
	}
	file, err := root.Open(filepath.Base(absolute))
	if err != nil {
		return nil, errors.Join(err, root.Close())
	}
	return &rootedFile{File: file, root: root}, nil
}

type rootedFile struct {
	*os.File
	root *os.Root
}

func (file *rootedFile) Close() error {
	fileErr := file.File.Close()
	rootErr := file.root.Close()
	if fileErr != nil {
		return fileErr
	}
	return rootErr
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeFileAtomically(data []byte, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	pattern := fmt.Sprintf("%s-*.tmp", filepath.Base(path))
	tempFile, err := os.CreateTemp(filepath.Dir(path), pattern)
	if err != nil {
		return err
	}
	defer func() {
		_ = tempFile.Close()
		_ = os.Remove(tempFile.Name())
	}()

	if err := os.Chmod(tempFile.Name(), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(tempFile.Name(), data, 0644); err != nil {
		return err
	}

	if err := tempFile.Close(); err != nil {
		return err
	}

	if err := os.Rename(tempFile.Name(), path); err != nil {
		return err
	}

	return nil
}

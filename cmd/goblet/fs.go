package main

import (
	"os"
	"path/filepath"
)

func writeFileAtomically(data []byte, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	pattern := filepath.Base(path) + "-*.tmp"
	tempFile, err := os.CreateTemp(filepath.Dir(path), pattern)
	if err != nil {
		return err
	}
	defer func() {
		_ = tempFile.Close()
		_ = os.Remove(tempFile.Name())
	}()

	if _, err := tempFile.Write(data); err != nil {
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

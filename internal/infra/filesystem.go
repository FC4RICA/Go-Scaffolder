package infra

import (
	"os"
	"path/filepath"
)

func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func WriteFile(path string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, content, 0644)
}

package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateProject(name string) error {
	if err := os.MkdirAll(name, 0755); err != nil {
		return err
	}

	// clean architecture layers
	dirs := []string{
		"cmd/" + name,
		"internal/core/domain",
		"internal/core/usecase",
		"internal/adapter/repository",
		"internal/adapter/handler",
		"pkg", // optional helpers
	}

	for _, d := range dirs {
		path := filepath.Join(name, d)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", path, err)
		}
	}

	// create main.go
	mainFile := filepath.Join(name, "cmd", name, "main.go")
	if err := os.WriteFile(mainFile, []byte(`package main

import "fmt"

func main() {
	fmt.Println("Hello from `+name+`!")
}
`), 0644); err != nil {
		return err
	}

	return nil
}

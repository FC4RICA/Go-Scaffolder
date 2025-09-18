package project

import (
	"fmt"
	"path/filepath"

	"github.com/FC4RICA/Go-Scaffolder/internal/infra"
)

func scaffoldBase(cfg Config) error {

	for _, d := range cfg.Dirs {
		d = filepath.Join(cfg.Path, d)
		if err := infra.CreateDir(d); err != nil {
			return fmt.Errorf("create dir %s: %w", d, err)
		}
	}

	return nil
}

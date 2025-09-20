package project

import "path/filepath"

func CreateProject(name string, path string) error {
	cfg := Config{
		ProjectName: name,
		Path:        filepath.Join(path, name),
		Dirs: []string{
			"cmd",
			"pkg",
			"internal/adapter/handler",
			"internal/adapter/repository",
			"internal/adapter/middleware",
			"internal/adapter/router",
			"internal/domain",
			"internal/interface",
			"internal/usecase",
		},
	}

	if err := scaffoldBase(cfg); err != nil {
		return err
	}

	return nil
}

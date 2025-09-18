package project

func CreateProject(name string) error {
	cfg := Config{
		ProjectName: name,
		Path:        "./" + name,
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

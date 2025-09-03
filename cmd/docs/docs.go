package docs

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed migration.txt
var migrationTemplate string

func InitDocs(projectName, modulePath string) error {
	configPath := filepath.Join(projectName, "docs")
	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(configPath, "migrations.md")
	if err := os.WriteFile(filePath, []byte(migrationTemplate), 0644); err != nil {
		return err
	}

	return nil
}

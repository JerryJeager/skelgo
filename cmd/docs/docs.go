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

	sourcePath := "./cmd/docs/migration.txt"

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	filePath := filepath.Join(configPath, "migrations.md")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(migrationTemplate)
	if err != nil {
		return err
	}

	return nil
}

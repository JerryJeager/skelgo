package docs

import (
	"io"
	"os"
	"path/filepath"
)

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

	_, err = io.Copy(file, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

package db

import (
	"os"
	"os/exec"
	"path/filepath"
)

func CreateMigrationFile(projectName string) error {
	projectPath := filepath.Join(projectName)

	migration := exec.Command("migrate", "create", "-ext", "psql", "-dir", "internal/db/migrations", "-seq", "create-users-table")
	migration.Dir = projectPath
	migration.Stdout = os.Stdout
	migration.Stderr = os.Stderr
	err := migration.Run()
	if err != nil {
		return err
	}
	return nil

}

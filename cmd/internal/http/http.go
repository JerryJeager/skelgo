package http

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed token.txt
var tokenTemplate string

//go:embed http.txt
var httpTemplate string

func HandleHttp(projectName, modulePath string) error {
	httpDirPath := filepath.Join(projectName, "internal", "http")
	err := os.MkdirAll(httpDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	tokenPath := filepath.Join(httpDirPath, "token.go")
	if err := os.WriteFile(tokenPath, []byte(tokenTemplate), 0644); err != nil {
		return err
	}

	httpPath := filepath.Join(httpDirPath, "http.go")
	if err := os.WriteFile(httpPath, []byte(httpTemplate), 0644); err != nil {
		return err
	}

	usersPath := filepath.Join(httpDirPath, "users.go")
	if err := os.WriteFile(usersPath, []byte(GenerateUserController(modulePath)), 0644); err != nil {
		return err
	}

	return nil
}

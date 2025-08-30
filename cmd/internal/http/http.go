package http

import (
	"io"
	"os"
	"path/filepath"
)

func HandleHttp(projectName, modulePath string) error {
	httpDirPath := filepath.Join(projectName, "internal", "http")
	err := os.MkdirAll(httpDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	tokenSource := "./cmd/internal/http/token.txt"
	httpSource := "./cmd/internal/http/http.txt"

	tokenFile, err := os.Open(tokenSource)
	if err != nil {
		return err
	}
	defer tokenFile.Close()

	httpFile, err := os.Open(httpSource)
	if err != nil {
		return err
	}
	defer httpFile.Close()

	tokenPath := filepath.Join(httpDirPath, "token.go")
	tokenFileDest, err := os.Create(tokenPath)
	if err != nil {
		return err
	}
	defer tokenFileDest.Close()
	_, err = io.Copy(tokenFileDest, tokenFile)
	if err != nil {
		return err
	}

	httpPath := filepath.Join(httpDirPath, "http.go")
	httpFileDest, err := os.Create(httpPath)
	if err != nil {
		return err
	}
	defer httpFileDest.Close()

	_, err = io.Copy(httpFileDest, httpFile)
	if err != nil {
		return err
	}

	usersPath := filepath.Join(httpDirPath, "users.go")
	usersFileDest, err := os.Create(usersPath)
	if err != nil {
		return err
	}
	defer usersFileDest.Close()

	_, err = usersFileDest.WriteString(GenerateUserController(modulePath))
	if err != nil {
		return err
	}

	return nil
}

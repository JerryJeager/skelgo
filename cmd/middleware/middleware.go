package middleware

import (
	"os"
	"path/filepath"
)

func HandleMiddleware(projectName, modulePath string) error {
	middlewarePath := filepath.Join(projectName, "middleware")
	err := os.MkdirAll(middlewarePath, os.ModePerm)
	if err != nil {
		return err
	}

	corsPath := filepath.Join(middlewarePath, "cors.go")
	corsFileDest, err := os.Create(corsPath)
	if err != nil{
		return err
	}
	defer corsFileDest.Close()

	_, err = corsFileDest.WriteString(GenerateCors())
	if err != nil{
		return err
	}

	authPath := filepath.Join(middlewarePath, "auth.go")
	authFileDest, err := os.Create(authPath)
	if err != nil{
		return err
	}
	defer authFileDest.Close()

	_, err = authFileDest.WriteString(GenerateAuthMiddleware(modulePath))
	if err != nil{
		return err
	}

	return nil
}

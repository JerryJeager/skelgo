package internal

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed template/models.txt
var modelsTemplate string

//go:embed template/otp.txt
var otpTemplate string

func CreateModels(projectName string) error {
	modelsPath := filepath.Join(projectName, "internal", "models")
	err := os.MkdirAll(modelsPath, os.ModePerm)
	if err != nil {
		return err
	}

	usersPath := filepath.Join(modelsPath, "users.go")
	if err := os.WriteFile(usersPath, []byte(modelsTemplate), 0644); err != nil {
		return err
	}

	otpPath := filepath.Join(modelsPath, "otp.go")
	if err := os.WriteFile(otpPath, []byte(otpTemplate), 0644); err != nil {
		return err
	}

	return nil
}

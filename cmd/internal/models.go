package internal

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/JerryJeager/skelgo/cmd/config"
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

	dependencies := []string{
		"github.com/google/uuid",
		"golang.org/x/crypto/bcrypt",
	}

	for _, dep := range dependencies {
		if err := config.DownloadDependency(filepath.Join(projectName), dep); err != nil {
			return err
		}
	}

	return nil
}

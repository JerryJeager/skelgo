package internal

import (
	"io"
	"os"
	"path/filepath"

	"github.com/JerryJeager/skelgo/cmd/config"
)

func CreateModels(projectName string) error {
	modelsPath := filepath.Join(projectName, "internal", "models")
	err := os.MkdirAll(modelsPath, os.ModePerm)
	if err != nil {
		return err
	}

	sourcePath := "./cmd/internal/template/models.txt"
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	otpSourcePath := "./cmd/internal/template/otp.txt"
	otpSource, err := os.Open(otpSourcePath)
	if err != nil {
		return err
	}
	defer otpSource.Close()

	usersPath := filepath.Join(modelsPath, "users.go")
	file, err := os.Create(usersPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, sourceFile)
	if err != nil {
		return err
	}

	otpPath := filepath.Join(modelsPath, "otp.go")
	otpFile, err := os.Create(otpPath)
	if err != nil {
		return err
	}
	defer otpFile.Close()

	_, err = io.Copy(otpFile, otpSource)
	if err != nil {
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

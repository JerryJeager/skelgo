package env

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func HandleEnvs(projectName string) error {
	projectPath := filepath.Join(projectName)
	envPath := filepath.Join(projectPath, ".env")
	envFileDest, err := os.Create(envPath)
	if err != nil {
		return err
	}
	defer envFileDest.Close()
	_, err = envFileDest.WriteString(getEnvTxt(projectName))
	if err != nil {
		return err
	}

	envExampleSource := "./cmd/internal/env/env.example.txt"
	envExampleFile, err := os.Open(envExampleSource)
	if err != nil {
		return err
	}
	defer envExampleFile.Close()

	envExamplePath := filepath.Join(projectPath, ".env.example")
	envExampleFileDest, err := os.Create(envExamplePath)
	if err != nil {
		return err
	}
	defer envExampleFileDest.Close()

	_, err = io.Copy(envExampleFileDest, envExampleFile)
	if err != nil {
		return err
	}
	return nil
}

func getEnvTxt(projectName string) string {
	return fmt.Sprintf(`
ENVIRONMENT=development
HOST=localhost
USER=postgres
PASSWORD="mydatabasepassword"
DBPORT=5432
DBNAME=%s
PORT=8080
JWT_SECRET="myjwtsecret"
JWT_EXPIRY="1"
EMAIL="myemail@gmail.com"
EMAILPASSWORD="figh qlfn hkbm xzqf"
EMAILUSERNAME="myemail@gmail.com"
	`, strings.ToLower(projectName))
}

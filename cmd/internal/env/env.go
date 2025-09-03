package env

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed env.example.txt
var envExampleTemplate string

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

	envExamplePath := filepath.Join(projectPath, ".env.example")
	if err := os.WriteFile(envExamplePath, []byte(envExampleTemplate), 0644); err != nil {
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

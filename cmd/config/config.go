package config

import (
	"os"
	"os/exec"
	"path/filepath"
)

func InitConfig(projectName, modulePath string) error {
	configPath := filepath.Join(projectName, "config")
	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(configPath, "config.go")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content := `package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func GetSession() *gorm.DB {
	return Session
}

func ConnectToDB() {
	environment := os.Getenv("ENVIRONMENT")
	var db *gorm.DB
	var err error
	if environment == "development" {
		//local development DB config:::
		host := os.Getenv("HOST")
		username := os.Getenv("USER")
		password := os.Getenv("PASSWORD")
		port := os.Getenv("DBPORT")
		dbName := os.Getenv("DBNAME")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		//production DB config:::
		connectionString := os.Getenv("CONNECTION_STRING")
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	}

	if err != nil {
		log.Fatal(err)
	}

	Session = db.Session(&gorm.Session{SkipDefaultTransaction: true, PrepareStmt: false})
	if Session != nil {
		log.Print("success: created db session")
	}
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Print(err)
		log.Print("failed to load envirionment variables")
		// log.Fatal("failed to load environment variables")
	}
}
`
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	dependencies := []string{
		"github.com/joho/godotenv",
		"gorm.io/driver/postgres",
		"gorm.io/gorm",
		"github.com/gin-gonic/gin",
	}

	for _, dep := range dependencies {
		if err := DownloadDependency(filepath.Join(projectName), dep); err != nil {
			return err
		}
	}

	return nil
}


func DownloadDependency(projectPath, dependency string) error {
	getCmd := exec.Command("go", "get", dependency)
	getCmd.Dir = projectPath
	getCmd.Stdout = os.Stdout
	getCmd.Stderr = os.Stderr
	err := getCmd.Run()
	if err != nil {
		return err
	}
	return nil
}
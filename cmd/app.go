package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/JerryJeager/skelgo/cmd/config"
)

func InitProject(projectName, modulePath string) error {
	projectPath := filepath.Join(projectName)
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	//create go mod
	cmd := exec.Command("go", "mod", "init", modulePath)
	cmd.Dir = projectPath 
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	//handle config folder
	if err := config.InitConfig(projectName, modulePath); err != nil {
		return err
	}

	return nil
}

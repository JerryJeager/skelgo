package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/JerryJeager/skelgo/cmd/config"
	"github.com/JerryJeager/skelgo/cmd/docs"
	"github.com/JerryJeager/skelgo/cmd/internal"
	"github.com/JerryJeager/skelgo/cmd/internal/utils"
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

	// handle config folder
	if err := config.InitConfig(projectName, modulePath); err != nil {
		return err
	}

	//handle docs folder 
	if err := docs.InitDocs(projectName, modulePath); err != nil {
		return err
	}

	//handle models 
	if err := internal.CreateModels(projectName); err != nil{
		return err
	}

	//handle utils 
	if err := utils.HandleUtils(projectName); err != nil{
		return err 
	}

	//tidy go modules
	tidy := exec.Command("go", "mod", "tidy")
	tidy.Dir = projectPath 
	tidy.Stdout = os.Stdout
	tidy.Stderr = os.Stderr
	err = tidy.Run()
	if err != nil {
		return err
	}

	return nil
}

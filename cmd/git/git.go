package git

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GitInitProject(projectName string) error {
	projectPath := filepath.Join(projectName)
	gitignorePath := filepath.Join(projectPath, ".gitignore")
	gitignoreDest, err := os.Create(gitignorePath)
	if err != nil {
		return err
	}
	defer gitignoreDest.Close()

	_, err = gitignoreDest.WriteString(getGitignoreTxt())
	if err != nil {
		return err
	}

	gitInit := exec.Command("git", "init")
	gitInit.Dir = projectPath
	gitInit.Stdout = os.Stdout
	gitInit.Stderr = os.Stderr
	err = gitInit.Run()
	if err != nil {
		return err
	}

	return nil
}

func getGitignoreTxt() string {
	return `
.vscode
.env
*.exe
	`
}

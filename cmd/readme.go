package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func HandleReadMe(projectName string) error {
	readmePath := filepath.Join(projectName, "README.md")
	readmeFileDest, err := os.Create(readmePath)
	if err != nil {
		return err
	}

	defer readmeFileDest.Close()

	mainReadme, err := os.ReadFile("./cmd/readme.txt")
	if err != nil {
		return err
	}
	txtReadMe := fmt.Sprintln(string(mainReadme))
	readmeContent := fmt.Sprintf(`
# %s
%s
	`, projectName, txtReadMe)

	_, err = readmeFileDest.WriteString(readmeContent)
	if err != nil {
		return err
	}

	return nil
}


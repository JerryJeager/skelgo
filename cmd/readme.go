package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed readme.txt
var readmeTemplate string

func HandleReadMe(projectName string) error {
	readmePath := filepath.Join(projectName, "README.md")
	readmeFileDest, err := os.Create(readmePath)
	if err != nil {
		return err
	}

	defer readmeFileDest.Close()

	readmeContent := fmt.Sprintf(`
# %s
%s
	`, projectName, readmeTemplate)

	_, err = readmeFileDest.WriteString(readmeContent)
	if err != nil {
		return err
	}

	return nil
}

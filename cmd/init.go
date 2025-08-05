package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <project-name> <module-path>",
	Short: "Initialize a new Go REST API project",
	Long: `The 'init' command generates a new Go project with a predefined 
folder structure and starter files using Gin, GORM, and PostgreSQL.

Provide the project name and the Go module path as arguments. 
Example: skelgo init myapp github.com/yourusername/myapp`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := Init(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Println("Project initialized successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

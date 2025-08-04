package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "skelgo",
	Short: "Scaffold a Go REST API project with Gin, GORM, and PostgreSQL",
	Long: `Skelgo is a personal CLI tool that generates a boilerplate Go project for building RESTful APIs. It sets up a standard folder structure, initializes a Go module, and includes starter code with Gin, GORM, and PostgreSQL integration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Skelgo is a scaffolding tool. Use a subcommand like 'skelgo init <project-name>'")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing skelgo '%s'\n", err)
		os.Exit(1)
	}
}

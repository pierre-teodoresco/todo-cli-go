/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const dbContextKey = "db"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A simple command-line todo application",
	Long: `todo-cli is a todo application for your terminal. Manage your tasks directly from the command line.

Examples:
  todo-cli add <title>
  todo-cli list
  todo-cli complete <id>
  todo-cli delete <id>

This CLI tool is built with Cobra, a powerful Go library for creating modern command-line applications.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Init DB
		db, err := sql.Open("sqlite3", os.Getenv("DB_URL"))
		if err != nil {
			log.Fatal("Failed to establish connection with database:", err)
		}
		ctx := context.WithValue(cmd.Context(), dbContextKey, db)
		cmd.SetContext(ctx)
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

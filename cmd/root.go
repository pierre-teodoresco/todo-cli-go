/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo-cli/internal/repository"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// App holds application dependencies
type App struct {
	DB      *sql.DB
	Queries *repository.Queries
}

// Global app instance
var app *App

// getDBPath returns the database path from env or default
func getDBPath() string {
	if dbURL := os.Getenv("DB_URL"); dbURL != "" {
		return dbURL
	}
	return "todo.db" // default path
}

// initApp initializes the application dependencies
func initApp() (*App, error) {
	dbPath := getDBPath()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	queries := repository.New(db)

	return &App{
		DB:      db,
		Queries: queries,
	}, nil
}

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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if app == nil {
			var err error
			app, err = initApp()
			if err != nil {
				log.Fatal("Failed to initialize application:", err)
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Cleanup resources
		if app != nil && app.DB != nil {
			app.DB.Close()
		}
	},
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

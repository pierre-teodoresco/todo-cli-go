/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"todo-cli/internal/handler"

	"github.com/spf13/cobra"
)

// addCmd represents the new command
var addCmd = &cobra.Command{
	Use:   "add [TITLE]",
	Short: "Add a new task to your list",
	Long:  `Adds a new task to your todo list. You must provide a title for the task. Example: todo-cli add "Buy groceries"`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]

		task, err := app.Queries.NewTask(cmd.Context(), title)
		if err != nil {
			log.Fatal("Error while creating a new task:", err)
		}
		handler.PrintCreatedTask(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

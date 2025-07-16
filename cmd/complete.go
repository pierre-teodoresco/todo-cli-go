/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"
	"todo-cli/internal/handler"
	"todo-cli/internal/repository"

	"github.com/spf13/cobra"
)

var reverseFlag bool

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [ID]",
	Short: "Mark a task as completed",
	Long: `Marks the specified task as completed.

Provide the task ID as an argument.
Use the -r or --reverse flag to mark a completed task as unfinished.

Examples:
  todo-cli complete 3
  todo-cli complete 3 --reverse
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Task ID should be a valid number:", err)
		}

		var task repository.Task
		if !reverseFlag {
			task, err = app.Queries.CompleteTask(cmd.Context(), int64(id))
		} else {
			task, err = app.Queries.ReverseCompleteTask(cmd.Context(), int64(id))
		}
		if err != nil {
			log.Fatal("Error while updating task:", err)
		}
		handler.PrintCompletedTask(task)
	},
}

func init() {
	completeCmd.Flags().BoolVarP(&reverseFlag, "reverse", "r", false, "Mark a task as unfinished")
	rootCmd.AddCommand(completeCmd)
}

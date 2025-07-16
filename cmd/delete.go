/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete a task by its ID",
	Long: `Deletes the specified task from your todo list.

Provide the task ID as an argument.

Example:
  todo-cli delete 5
This will remove the task with ID 5 from your list.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Task ID should be a valid number:", err)
		}

		err = app.Queries.DeleteTask(cmd.Context(), int64(id))
		if err != nil {
			log.Fatal("Error while deleting task:", err)
		}
		fmt.Printf("Task %d has been deleted\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

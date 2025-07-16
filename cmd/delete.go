/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"todo-cli/internal/repository"

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

		ctx := cmd.Context()
		db := ctx.Value(dbContextKey).(*sql.DB)
		queries := repository.New(db)

		err = queries.DeleteTask(ctx, int64(id))
		if err != nil {
			log.Fatal("Error while creating a new task:", err)
		}
		fmt.Printf("Task %d has been deleted\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

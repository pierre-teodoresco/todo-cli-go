/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"log"
	"todo-cli/internal/handler"
	"todo-cli/internal/repository"

	"github.com/spf13/cobra"
)

var allFlag bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [-a|--all]",
	Short: "Show unfinished tasks",
	Long:  `Displays the list of unfinished tasks. Use -a or --all to show all tasks, including completed ones.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		db := ctx.Value(dbContextKey).(*sql.DB)
		queries := repository.New(db)

		var tasks []repository.Task
		var err error
		if allFlag {
			tasks, err = queries.FindAllTasks(ctx)
		} else {
			tasks, err = queries.FindAllUnfinishedTasks(ctx)
		}
		if err != nil {
			log.Fatal("Error while creating a new task:", err)
		}
		handler.PrintTasks(tasks)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "List all tasks even the finished ones")
	rootCmd.AddCommand(listCmd)
}

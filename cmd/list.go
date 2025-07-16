package cmd

import (
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
		var tasks []repository.Task
		var err error
		if allFlag {
			tasks, err = app.Queries.FindAllTasks(cmd.Context())
		} else {
			tasks, err = app.Queries.FindAllUnfinishedTasks(cmd.Context())
		}
		if err != nil {
			log.Fatal("Error while fetching tasks:", err)
		}
		handler.PrintTasks(tasks)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "List all tasks even the finished ones")
	rootCmd.AddCommand(listCmd)
}

package handler

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
	"todo-cli/internal/repository"

	"github.com/mergestat/timediff"
)

func nullTimeString(t sql.NullTime) string {
	if t.Valid {
		return timediff.TimeDiff(t.Time.Add(10), timediff.WithStartTime(time.Now()))
	}
	return "Error while parsing time"
}

func int64BoolString(b int64) string {
	var done string
	if b != 0 {
		done = "true"
	} else {
		done = "false"
	}
	return done
}

func PrintCreatedTask(task repository.Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	id := strconv.Itoa(int(task.ID))
	title := task.Title
	createdAt := nullTimeString(task.CreatedAt)
	fmt.Fprintln(w, "ID\tTask\tCreated")
	fmt.Fprintf(w, "%s\t%s\t%s\n", id, title, createdAt)
	w.Flush()
}

func PrintTasks(tasks []repository.Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
	for _, task := range tasks {
		id := strconv.Itoa(int(task.ID))
		title := task.Title
		createdAt := nullTimeString(task.CreatedAt)
		done := int64BoolString(task.Completed)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", id, title, createdAt, done)
	}
	w.Flush()
}

func PrintCompletedTask(task repository.Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	id := strconv.Itoa(int(task.ID))
	title := task.Title
	createdAt := nullTimeString(task.CreatedAt)
	done := int64BoolString(task.Completed)
	fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", id, title, createdAt, done)
	w.Flush()
}

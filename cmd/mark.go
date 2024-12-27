package cmd

import (
	"log"
	"strconv"

	"github.com/Metudu/task-tracker/task"
	"github.com/spf13/cobra"
)

var markInProgress = &cobra.Command{
	Use: "mark-in-progress",
	Short: "mars a task as in progress :P",
	Args: cobra.ExactArgs(1),
	Run: markInProgressFunc,
}

var markInProgressFunc = func(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	task.MarkAsInProgress(id)
}

var markDone = &cobra.Command{
	Use: "mark-done",
	Short: "mars a task as in progress :P",
	Args: cobra.ExactArgs(1),
	Run: markDoneFunc,
}

var markDoneFunc = func(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	task.MarkAsDone(id)
}
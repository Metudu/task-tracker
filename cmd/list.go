package cmd

import (
	"log"

	"github.com/Metudu/task-tracker/task"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use: "list",
	Short: "list task(s) :P",
	Args: cobra.MaximumNArgs(1),
	Run: listFunc,
}

var listFunc = func(cmd *cobra.Command, args []string) {
	if len(args) == 1 && args[0] != "done" && args[0] != "in-progress" && args[0] != "todo" {
		log.Fatalf("Invalid argument: %v\n", args[0])
	}
	if len(args) == 0 {
		task.ListTasks("")
		return
	}
	task.ListTasks(args[0])
}
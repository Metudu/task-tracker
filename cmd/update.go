package cmd

import (
	"log"
	"strconv"

	"github.com/Metudu/task-tracker/task"
	"github.com/spf13/cobra"
)

var update = &cobra.Command{
	Use: "update",
	Short: "updates a task :P",
	Args: cobra.ExactArgs(2),
	Run: updateFunc,
}

var updateFunc = func(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Argument %v is not an integer value!\n", args[0])
	}

	task.UpdateTask(id, args[1])
}
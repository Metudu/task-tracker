package cmd

import (
	"log"
	"strconv"

	"github.com/Metudu/task-tracker/task"
	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use: "delete",
	Short: "deletes a task :P",
	Args: cobra.ExactArgs(1),
	Run: deleteFunc,
}

var deleteFunc = func(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Argument %v is not an integer value!\n", args[0])
	}

	task.DeleteTask(id)
}
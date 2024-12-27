package cmd

import (
	"github.com/Metudu/task-tracker/task"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use: "add",
	Short: "adds a task :P",
	Args: cobra.ExactArgs(1),
	Run: addFunc,
}


var addFunc = func(cmd *cobra.Command, args []string) {
	task.AddTask(args[0])
}
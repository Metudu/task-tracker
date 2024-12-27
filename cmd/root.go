package cmd

import (
	"github.com/spf13/cobra"
)

func init(){
	root.AddCommand(add, update, delete, markInProgress, markDone, list)
}

var root = &cobra.Command{
	Use: "task-tracker",
	Short: "Tracks your tasks :P",
}

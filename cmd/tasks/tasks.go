package tasks

import (
	"github.com/spf13/cobra"

	"util/cmd/tasks/cleanup"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "tasks",
	Long:  `tasks`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	cmd.Help()
	// },
}

func init() {
	TaskCmd.AddCommand(cleanup.CleanupCmd)
}

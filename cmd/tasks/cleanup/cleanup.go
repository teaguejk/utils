package cleanup

import (
	"github.com/spf13/cobra"

	"util/cmd/tasks/cleanup/dir"
)

var CleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "cleanup task",
	Long:  `cleanup task`,
}

func init() {
	CleanupCmd.AddCommand(dir.CleanupDirCmd)
}

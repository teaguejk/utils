package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var CleanupDirCmd = &cobra.Command{
	Use:   "dir",
	Short: "cleanup dir",
	Long:  `cleanup dir`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := CleanupDir(); err != nil {
			fmt.Println(err)
		}
	},
}

var path string
var duration string

func init() {
	CleanupDirCmd.Flags().StringVarP(&path, "path", "p", "", "path to directory")
	CleanupDirCmd.Flags().StringVarP(&duration, "time", "t", "", "the minimum duration of a file to delete")
	CleanupDirCmd.MarkFlagRequired("path")
}

func CleanupDir() error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if duration == "" {
			fmt.Println(path)

			err := os.Remove(path)
			if err != nil {
				fmt.Printf("error while deleting file: %s\n", path)
			} else {
				fmt.Printf("deleted file: %s\n", path)
			}

			return nil
		}

		t, err := time.ParseDuration(duration)
		if err != nil {
			return err
		}

		if time.Since(info.ModTime()) > t {
			fmt.Println(path, info.ModTime())

			err := os.Remove(path)
			if err != nil {
				fmt.Printf("error while deleting file: %s\n", path)
			} else {
				fmt.Printf("deleted file: %s\n", path)
			}
		}

		return nil
	})

	return err
}

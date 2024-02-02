package cli

import (
	"fmt"
	"os"

	"github.com/dustinliu/taskcommander/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to inbox",
	Long:  `Add a task to inbox`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := service.NewService()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

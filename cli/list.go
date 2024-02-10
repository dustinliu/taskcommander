package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Run: func(_ *cobra.Command, _ []string) {
		service := getService()

		tasks, err := service.ListTodoTasks()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, task := range tasks {
			fmt.Printf("%+v\n", task)
		}
	},
}

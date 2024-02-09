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
	Run: func(_ *cobra.Command, argv []string) {
		srv := getService()

		task := srv.NewTask().SetTitle(argv[0]).SetCategory(service.CategorySomeday)
		task, err := srv.AddTask(task)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", task)
	},
}

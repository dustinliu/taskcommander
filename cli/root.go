package cli

import (
	"fmt"
	"os"

	"github.com/dustinliu/taskcommander/core"
	"github.com/spf13/cobra"
)

const (
	AppName = "TaskCommander"
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&core.Debug, "debug", "d", false, "debug mode")
}

var rootCmd = &cobra.Command{
	Use:   "tc",
	Short: "taskcommand cli tool",
	//Run: func(cmd *cobra.Command, args []string) {
	//// Do Stuff Here
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cli

import (
	"fmt"
	"os"
	"sync"

	"github.com/dustinliu/taskcommander/core"
	"github.com/dustinliu/taskcommander/service"
	"github.com/spf13/cobra"
)

var (
	srv  service.TaskService
	once sync.Once
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&core.Debug, "debug", "d", false, "debug mode")
}

var rootCmd = &cobra.Command{
	Use:   "tc",
	Short: "taskcommand cli tool",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getService() service.TaskService {
	once.Do(func() {
		var err error
		srv, err = service.NewService()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := srv.OAuth2(showUrl); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	})
	return srv
}

func showUrl(url string) error {
	fmt.Printf("Please visit the following URL to authorize this application:\n%s\n", url)
	return nil
}

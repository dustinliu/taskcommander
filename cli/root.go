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

func getService() service.TaskService {
	once.Do(func() {
		var err error
		srv, err = service.NewService()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if srv.InitOauth2Needed() {
			authUrl := srv.GetOauthAuthUrl()
			fmt.Printf("Please visit the following URL to authorize this application:\n%s\n", authUrl)
			if err := srv.WaitForAuthDone(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
		if err := srv.Init(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	})
	return srv
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&core.Debug, "debug", "d", false, "debug mode")
	cobra.OnInitialize(func() {
		service, err := service.NewService()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if service.InitOauth2Needed() {
			authUrl := service.GetOauthAuthUrl()
			fmt.Printf("Please visit the following URL to authorize this application:\n%s\n", authUrl)
			if err := service.WaitForAuthDone(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	})
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

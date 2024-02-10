package cli

import (
	"fmt"
	"os"

	"github.com/dustinliu/taskcommander/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth to google task",
	Run: func(_ *cobra.Command, _ []string) {
		service, err := service.NewGoogleTaskService()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create service: %v\n", err)
			os.Exit(1)
		}
		if service.InitOauth2Needed() {
			authURL := service.GetOauthAuthUrl()
			fmt.Printf("Go to the following link in your browser: \n%v\n", authURL)

			err := service.WaitForAuthDone()
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to fetch oauth token: %v", err)
				os.Exit(1)
			}
		}
	},
}

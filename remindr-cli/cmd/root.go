package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "remindr",
	Short: "Remindr is a recurring reminder service written for Linux users",
	Long:  "Remindr is a recurring reminder service written for Linux users in Golang. It works similar to Docker where you have to setup a systemd user service for the server and interact with reminders using CLI commands.",
	Run: func(cmd *cobra.Command, args []string) {
    cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

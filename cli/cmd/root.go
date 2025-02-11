package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "remindr",
	Short: "Remindr is a reminder service written in Golang for Linux",
	Long:  "Baadme add kardena",
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

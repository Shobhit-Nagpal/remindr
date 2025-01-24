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
		fmt.Println("Yo, just running that root, you know what i mean")
		fmt.Println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

	jobsCmd.AddCommand(listCmd)
	jobsCmd.AddCommand(killCmd)
	jobsCmd.AddCommand(stopCmd)

	rootCmd.AddCommand(jobsCmd)
}

var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Sub command for jobs",
	Long:  `Fill this later`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Jobs command basically")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Sub command for jobs",
	Long:  `Fill this later`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List deez")
	},
}

var killCmd = &cobra.Command{
	Use: "kill",
	//Acts as a custom validator
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Kill command required id of job")
		}

		return nil
	},
	Short: "Sub command for jobs",
	Long:  `Fill this later`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Killing child job id ", args[0])
	},
}

var stopCmd = &cobra.Command{
	Use: "stop",
	//Acts as a custom validator
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Stop command required id of job")
		}

		return nil
	},
	Short: "Sub command for jobs",
	Long:  `Fill this later`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping child job id ", args[0])
	},
}

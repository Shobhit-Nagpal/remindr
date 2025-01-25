package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type JobPayload struct {
	ID        uuid.UUID     `json:"id"`
	Message   string        `json:"message"`
	Interval  time.Duration `json:"interval"`
	Level     jobs.Level         `json:"level"`
	Active    bool          `json:"active"`
	CreatedAt time.Time     `json:"created_at"`
}

func init() {
	// Add flags
	listCmd.Flags().Bool("all", false, "Show all reminders, including inactive ones")

	createCmd.Flags().Float32("interval", 600, "Interval after which reminder notification pops up (in seconds)")
	createCmd.Flags().String("level", "normal", "Urgency level of the reminder notification")

	//Add sub comands
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(killCmd)
	rootCmd.AddCommand(stopCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Sub command for listing reminders",
	Long:  `Fill this later`,
	Run: func(cmd *cobra.Command, args []string) {
    res, err := http.Get("http://localhost:5678/api/reminders")
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Fatalf("Error getting jobs\n")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		jobs := []*JobPayload{}

		err = json.Unmarshal(body, &jobs)
		if err != nil {
			log.Fatal(err)
		}

    for _, job := range jobs {
      fmt.Println(job)
    }

	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Sub command for jobs",
	Long:  `Fill this later`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Create command requires a message to create a reminder")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Message of create command", args[0])
		fmt.Println("Create deez")
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

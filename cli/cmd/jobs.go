package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/spf13/cobra"
)

type JobPayload struct {
	ID        string  `json:"id"`
	Message   string     `json:"message"`
	Interval  int        `json:"interval"`
	Level     jobs.Level `json:"level"`
	Active    bool       `json:"active"`
	CreatedAt time.Time  `json:"created_at"`
}

func init() {
	// Add flags
	listCmd.Flags().Bool("all", false, "Show all reminders, including inactive ones")

	createCmd.Flags().Int("interval", 6, "Interval after which reminder notification pops up (in seconds)")
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

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		jobs := []*JobPayload{}

		err = json.Unmarshal(body, &jobs)
		if err != nil {
			log.Fatal(err)
		}

		if len(jobs) == 0 {
			fmt.Println("No jobs running")
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
		message := args[0]

		level, err := cmd.Flags().GetString("level")
		if err != nil {
			log.Fatal(err)
		}

		interval, err := cmd.Flags().GetInt("interval")
		if err != nil {
			log.Fatal(err)
		}

		job := jobs.CreateJob(message, interval, jobs.Level(level))

		payload := JobPayload{
			ID:        job.ID,
			Message:   job.Message,
			Interval:  interval,
			Level:     job.Level,
			Active:    job.Active,
			CreatedAt: job.CreatedAt,
		}

		jobData, err := json.Marshal(payload)

		res, err := http.Post("http://localhost:5678/api/reminders", "application/json", bytes.NewBuffer(jobData))
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode == http.StatusCreated {
			fmt.Println("Job created and active")
		}
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

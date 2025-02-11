package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/Shobhit-Nagpal/remindr/internal/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type JobPayload struct {
	ID        string     `json:"id"`
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
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(destroyCmd)
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

		data := [][]string{}

		for _, job := range jobs {
			row := []string{strings.Split(job.ID, "-")[0], job.Message, string(job.Level), strconv.Itoa(job.Interval), strconv.FormatBool(job.Active), job.CreatedAt.String()}
			data = append(data, row)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Message", "Level", "Interval", "Active", "Created At"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render() // Send output
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
		if err != nil {
			log.Fatal(err)
		}

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
		id := args[0]

		type Payload struct {
			ID string `json:"id"`
		}

		payload := Payload{
			ID: id,
		}

		jobData, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("DELETE", "http://localhost:5678/api/reminders", bytes.NewBuffer(jobData))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode == http.StatusNoContent {
			fmt.Println("Job killed")
		}
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
		id := args[0]

		type Payload struct {
			ID     string `json:"id"`
			Active bool   `json:"active"`
		}

		payload := Payload{
			ID:     id,
			Active: false,
		}

		jobData, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("PUT", "http://localhost:5678/api/reminders", bytes.NewBuffer(jobData))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			fmt.Println("Job stopped")
		}
	},
}

var runCmd = &cobra.Command{
	Use: "run",
	//Acts as a custom validator
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Run command required id of job")
		}

		return nil
	},
	Short: "Sub command for jobs",
	Long:  `Fill this later`,
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		type Payload struct {
			ID     string `json:"id"`
			Active bool   `json:"active"`
		}

		payload := Payload{
			ID:     id,
			Active: true,
		}

		jobData, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("PUT", "http://localhost:5678/api/reminders", bytes.NewBuffer(jobData))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			fmt.Println("Job running")
		}
	},
}

var setupCmd = &cobra.Command{
	Use:   "setup [working-directory]",
	Short: "Setup Remindr as a service",
	Long:  `Setup Remindr as a systemd user service. Requires the working directory path as an argument.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("working directory path is required")
		}

		// Verify if the directory exists
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			return fmt.Errorf("working directory does not exist: %s", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		workingDir := args[0]

		// Get home directory
		homeDir, err := utils.GetHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		// Create systemd user directory if it doesn't exist
		systemdDir := fmt.Sprintf("%s/.config/systemd/user", homeDir)
		if err := os.MkdirAll(systemdDir, 0755); err != nil {
			log.Fatal(err)
		}

		// Create service file
		remindrServiceFile := fmt.Sprintf("%s/remindr.service", systemdDir)
		file, err := os.Create(remindrServiceFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		content := getServiceFileContent(workingDir)
		if _, err := file.WriteString(content); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Service file created at: %s\n", remindrServiceFile)
		fmt.Println("To start the service, run:")
		fmt.Println("systemctl --user daemon-reload")
		fmt.Println("systemctl --user enable --now remindr.service")
	},
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Remove Remindr service",
	Long:  `Remove and cleanup the Remindr systemd service file`,
	Run: func(cmd *cobra.Command, args []string) {
		// First stop and disable the service
		stopCmd := exec.Command("systemctl", "--user", "stop", "remindr.service")
		if err := stopCmd.Run(); err != nil {
			fmt.Println("Warning: Could not stop service:", err)
		}

		disableCmd := exec.Command("systemctl", "--user", "disable", "remindr.service")
		if err := disableCmd.Run(); err != nil {
			fmt.Println("Warning: Could not disable service:", err)
		}

		// Get home directory
		homeDir, err := utils.GetHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		// Remove service file
		serviceFile := fmt.Sprintf("%s/.config/systemd/user/remindr.service", homeDir)
		if err := os.Remove(serviceFile); err != nil {
			if !os.IsNotExist(err) {
				log.Fatal("Error removing service file:", err)
			}
			fmt.Println("Service file not found")
			return
		}

		// Reload systemd daemon
		reloadCmd := exec.Command("systemctl", "--user", "daemon-reload")
		if err := reloadCmd.Run(); err != nil {
			fmt.Println("Warning: Could not reload systemd:", err)
		}

		fmt.Println("Remindr service has been removed successfully")
	},
}

func getServiceFileContent(workingDir string) string {
	return fmt.Sprintf(`[Unit]
Description=Remindr Service
After=default.target

[Service]
ExecStart=/usr/bin/go run server/main.go
WorkingDirectory=%s
Restart=on-failure

[Install]
WantedBy=default.target
`, workingDir)
}

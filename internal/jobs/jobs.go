package jobs

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

type Level string

const (
	LOW      Level = "low"
	NORMAL   Level = "normal"
	CRITICAL Level = "critical"
)

type Job struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Interval  int       `json:"interval"`
	Level     Level     `json:"level"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

// Functions

func CreateJob(message string, interval int, level Level) *Job {
	id := uuid.New()

	return &Job{
		ID:        id.String(),
		Message:   message,
		Interval:  interval,
		Level:     level,
		Active:    true,
		CreatedAt: time.Now(),
	}
}

// Methods

func (j *Job) SetMessage(message string) {
	j.Message = message
}

func (j *Job) SetInterval(interval int) {
	j.Interval = interval
}

func (j *Job) SetLevel(level Level) {
	switch level {
	case LOW, NORMAL, CRITICAL:
		j.Level = level
	default:
		j.Level = NORMAL
	}
}

func (j *Job) SetActive(active bool) {
	j.Active = active
}

func (j *Job) Notify() error {
	cmd := exec.Command("notify-send", "-u", fmt.Sprintf("%s", j.Level), "-t", "5000", j.Message)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

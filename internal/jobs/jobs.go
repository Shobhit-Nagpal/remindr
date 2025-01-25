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
  id        uuid.UUID
	message   string
	interval  time.Duration
	level     Level
	active    bool
	createdAt time.Time
}

// Functions

func CreateJob(message string, interval float32, level Level) *Job {
	id := uuid.New()

	return &Job{
		id:        id,
		message:   message,
		interval:  time.Duration(interval) * time.Second,
		level:     level,
		active:    true,
		createdAt: time.Now(),
	}
}


// Methods

func (j *Job) ID() uuid.UUID {
	return j.id
}

func (j *Job) Message() string {
	return j.message
}

func (j *Job) SetMessage(message string) {
	j.message = message
}

func (j *Job) Interval() time.Duration {
	return j.interval
}

func (j *Job) SetInterval(interval float32) {
	j.interval = time.Duration(interval) * time.Second
}

func (j *Job) Level() Level {
	return j.level
}

func (j *Job) SetLevel(level Level) {
	switch level {
	case LOW, NORMAL, CRITICAL:
		j.level = level
	default:
		j.level = NORMAL
	}
}

func (j *Job) Active() bool {
	return j.active
}

func (j *Job) SetActive(active bool) {
	j.active = active
}

func (j *Job) Notify() error {
	cmd := exec.Command("notify-send", "-u", fmt.Sprintf("%s", j.Level()), "-t", "5000", j.Message())
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (j *Job) CreatedAt() time.Time {
	return j.createdAt
}

package jobs

import (
	"fmt"
	"log"
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

func ScheduleJob(job *Job) *time.Ticker {
	ticker := time.NewTicker(job.Interval())
	return ticker
}

func RunAll(jobs []*Job) {
	for _, job := range jobs {
		ticker := ScheduleJob(job)

		fmt.Printf("\nScheduled job %s, for %s interval\n", job.ID(), job.Interval().String())

		go func() {
			for {
				select {
				case <-ticker.C:
					cmd := exec.Command("notify-send", "-u", fmt.Sprintf("%s", job.Level()), "-t", "5000", job.Message())
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
				}
			}
		}()
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

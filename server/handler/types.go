package handler

import (
	"time"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
)

type JobPayload struct {
	ID        string     `json:"id"`
	Message   string     `json:"message"`
	Interval  int        `json:"interval"`
	Level     jobs.Level `json:"level"`
	Active    bool       `json:"active"`
	CreatedAt time.Time  `json:"created_at"`
}

package handler

import (
	"time"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/google/uuid"
)

type JobPayload struct {
	ID        uuid.UUID     `json:"id"`
	Message   string        `json:"message"`
	Interval  time.Duration `json:"interval"`
	Level     jobs.Level         `json:"level"`
	Active    bool          `json:"active"`
	CreatedAt time.Time     `json:"created_at"`
}

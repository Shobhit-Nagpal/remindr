package jobs

import (
	"fmt"

	"github.com/google/uuid"
)

type JobManager struct {
	jobs map[uuid.UUID]*Job
}

func (jm *JobManager) RegisterJob(job *Job) {
	if _, exists := jm.jobs[job.ID()]; exists {
		fmt.Printf("Job already registered")
		return
	}

	jm.jobs[job.ID()] = job
}

func (jm *JobManager) ListAllJobs() {
}

func (jm *JobManager) ListActiveJobs() {
  fmt.Println("Active jobs: ")
  //Print a table here
}

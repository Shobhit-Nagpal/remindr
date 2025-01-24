package jobs

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type JobManager struct {
	jobs map[uuid.UUID]*Job
}

func CreateJobManager() *JobManager {
	return &JobManager{
		jobs: map[uuid.UUID]*Job{},
	}
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

func (jm *JobManager) ScheduleJob(id uuid.UUID) *time.Ticker {
	ticker := time.NewTicker(jm.jobs[id].Interval())
	return ticker
}

func (jm *JobManager) RunAllJobs() {
	for _, job := range jm.jobs {
		ticker := jm.ScheduleJob(job.ID())

		fmt.Printf("\nScheduled job %s, for %s interval\n", job.ID(), job.Interval().String())
		go func() {
			for {
				select {
				case <-ticker.C:
					err := job.Notify()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}()
	}
}

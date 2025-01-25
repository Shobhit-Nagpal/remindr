package jobs

import (
	"fmt"
	"log"
	"time"
)

type JobManager struct {
	jobs map[string]*Job
}

func CreateJobManager() *JobManager {
	return &JobManager{
		jobs: map[string]*Job{},
	}
}

func (jm *JobManager) RegisterJob(job *Job) {
	if _, exists := jm.jobs[job.ID]; exists {
		fmt.Printf("Job already registered")
		return
	}

	jm.jobs[job.ID] = job
}

func (jm *JobManager) RegisterJobs(jobs []*Job) {
	for _, job := range jobs {
		jm.RegisterJob(job)
	}
}

func (jm *JobManager) GetAllJobs() map[string]*Job {
	jobs := []*Job{}
	for _, job := range jm.jobs {
		jobs = append(jobs, job)
	}

	return jm.jobs
}

func (jm *JobManager) ListActiveJobs() {
	fmt.Println("Active jobs: ")
	//Print a table here
}

func (jm *JobManager) ScheduleJob(id string) *time.Ticker {
	ticker := time.NewTicker(time.Duration(jm.jobs[id].Interval) * time.Second)
	return ticker
}

func (jm *JobManager) RunAllJobs() {
	for _, job := range jm.jobs {

		if !job.Active {
			continue
		}
		jm.RunJob(job)
	}
}

func (jm *JobManager) RunJob(job *Job) {
	ticker := jm.ScheduleJob(job.ID)

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

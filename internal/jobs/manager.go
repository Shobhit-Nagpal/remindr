package jobs

import (
	"fmt"
	"log"
	"time"
)

type JobManager struct {
	jobs map[string]*Job
	stop chan string
}

func CreateJobManager() *JobManager {
	return &JobManager{
		jobs: map[string]*Job{},
		stop: make(chan string, 50),
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

func (jm *JobManager) UnregisterJob(job *Job) {
	if _, exists := jm.jobs[job.ID]; !exists {
		fmt.Printf("Job doesn't exist")
		return
	}

	delete(jm.jobs, job.ID)
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
	if !job.Active {
		return
	}

	jm.runInternalJob(job)
}

func (jm *JobManager) StartJob(job *Job) {
	jm.jobs[job.ID].Active = true

	jm.runInternalJob(job)
}

func (jm *JobManager) StopJob(job *Job) {
	if _, exists := jm.jobs[job.ID]; !exists {
		fmt.Printf("Job %s doesn't exist\n", job.ID)
		return
	}

	if jm.jobs[job.ID].Active {
		jm.stop <- job.ID
		fmt.Printf("Job %s stopped successfully\n", job.ID)
	} else {
		fmt.Printf("Job %s is not running\n", job.ID)
	}
}

func (jm *JobManager) runInternalJob(job *Job) {

	ticker := jm.ScheduleJob(job.ID)

	go func() {

		for {
			select {
			case <-ticker.C:
				err := job.Notify()
				if err != nil {
					log.Fatal(err)
				}
			case id := <-jm.stop:
				if id == job.ID {
					ticker.Stop()
					return
				}
			}
		}
	}()

}

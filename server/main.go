package main

import (
	"log"

	"github.com/Shobhit-Nagpal/remindr/internal/db"
	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/Shobhit-Nagpal/remindr/server/api"
)

func main() {
	// Initialize local db / dotfolder
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	jobsManager := jobs.CreateJobManager()

	// Git data
	allJobs, err := db.GetAllJobs()

	//Register jobs
	jobsManager.RegisterJobs(allJobs)

	job := jobs.CreateJob("yooooo", 6, "normal")
	job2 := jobs.CreateJob("namaste", 10, "critical")

	jobsManager.RegisterJob(job)
	jobsManager.RegisterJob(job2)

	//Spin up jobs
	jobsManager.RunAllJobs()

	server := api.NewServer(jobsManager)

	server.ListenAndServe()
}

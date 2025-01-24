package main

import (
	"log"
	"time"

	"github.com/Shobhit-Nagpal/remindr/internal/db"
	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
)

func main() {
	// Initialize local db / dotfolder
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	jobsManager := jobs.CreateJobManager()

	// Git data
	j1 := jobs.CreateJob("git up", 5, "critical")
	j2 := jobs.CreateJob("soja", 10, "normal")

	//Register jobs
	jobsManager.RegisterJob(j1)
	jobsManager.RegisterJob(j2)

	//Spin up jobs
	jobsManager.RunAllJobs()

	for {
		time.Sleep(time.Second)
	}
}

package main

import (
	"log"

	"github.com/Shobhit-Nagpal/remindr/internal/db"
	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/Shobhit-Nagpal/remindr/internal/utils"
	"github.com/Shobhit-Nagpal/remindr/server/api"
)

func main() {
	// Initialize local db / dotfolder
  dir, err := utils.GetDBPath()
	if err != nil {
		log.Fatal(err)
	}

  database, err := db.NewDB(dir, "db.json")
	if err != nil {
		log.Fatal(err)
	}

	jobsManager := jobs.CreateJobManager()

	// Git data
	reminders, err := database.GetAllJobs()

  allJobs := []*jobs.Job{}

  for _, reminder := range reminders {
    allJobs = append(allJobs, &reminder)
  }

	//Register jobs
	jobsManager.RegisterJobs(allJobs)

	//Spin up jobs
	jobsManager.RunAllJobs()

	server := api.NewServer(database, jobsManager)

	server.ListenAndServe()
}

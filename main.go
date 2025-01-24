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
	// Git data
	j1 := jobs.CreateJob("git up", 5, "critical")
	j2 := jobs.CreateJob("soja", 10, "normal")

	//Start process
	jobs.RunAll([]*jobs.Job{j1, j2})

	for {
		time.Sleep(time.Second)
	}
}

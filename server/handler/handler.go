package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-9")
	io.WriteString(w, "OK")
	w.WriteHeader(http.StatusOK)
}

func Index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-9")
	io.WriteString(w, "Welcome to BadgeFlow's backend")
	w.WriteHeader(http.StatusOK)
}

func GetReminders(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	if db == nil {
		respondWithError(w, http.StatusInternalServerError, "DB not found")
		return
	}

	reminders, err := db.GetAllJobs()
	if err != nil {
		log.Print(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't read jobs")
		return
	}

	response := []*JobPayload{}

	for _, reminder := range reminders {
		response = append(response, &JobPayload{
			ID:        reminder.ID,
			Message:   reminder.Message,
			Interval:  int(reminder.Interval),
			Level:     reminder.Level,
			Active:    reminder.Active,
			CreatedAt: reminder.CreatedAt,
		})
	}

	respondWithJSON(w, http.StatusOK, response)

}

func CreateReminder(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	if db == nil {
		respondWithError(w, http.StatusInternalServerError, "DB not found")
		return
	}

	manager := getManager(req)
	if manager == nil {
		respondWithError(w, http.StatusInternalServerError, "Manager not found")
		return
	}

	reminder := JobPayload{}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't read body")
		return
	}
	defer req.Body.Close()

	err = json.Unmarshal(body, &reminder)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't unmarshal reminder from body")
		return
	}

	job, err := db.CreateJob(reminder.Message, reminder.Interval, string(reminder.Level))
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create reminder")
		return
	}

	manager.RegisterJob(job)
	manager.RunJob(job)

	response := JobPayload{
		ID:        job.ID,
		Message:   job.Message,
		Interval:  int(job.Interval),
		Level:     job.Level,
		Active:    job.Active,
		CreatedAt: job.CreatedAt,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

func DeleteReminder(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	if db == nil {
		respondWithError(w, http.StatusInternalServerError, "DB not found")
		return
	}

	manager := getManager(req)
	if manager == nil {
		respondWithError(w, http.StatusInternalServerError, "Manager not found")
		return
	}

	type JobIDPayload struct {
		ID string `json:"id"`
	}

	jobPayload := JobIDPayload{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't read body")
		return
	}
	defer req.Body.Close()

	err = json.Unmarshal(body, &jobPayload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't unmarshal reminder from body")
		return
	}

	job, err := db.DeleteJob(jobPayload.ID)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create reminder")
		return
	}

	manager.StopJob(job)
	manager.UnregisterJob(job)

	response := JobPayload{
		ID:        job.ID,
		Message:   job.Message,
		Interval:  int(job.Interval),
		Level:     job.Level,
		Active:    job.Active,
		CreatedAt: job.CreatedAt,
	}

	respondWithJSON(w, http.StatusNoContent, response)
}

package handler

import (
	"io"
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
	manager := getManager(req)
	if manager == nil {
		respondWithError(w, http.StatusInternalServerError, "Manager not found")
		return
	}

	reminders := manager.GetAllJobs()

	response := []*JobPayload{}

	for _, reminder := range reminders {
		response = append(response, &JobPayload{
			ID:       reminder.ID(),
			Message:  reminder.Message(),
			Interval: reminder.Interval(),
			Level:    reminder.Level(),
      Active: reminder.Active(),
      CreatedAt: reminder.CreatedAt(),
		})
	}

	respondWithJSON(w, http.StatusOK, response)

}

func CreateReminder(w http.ResponseWriter, req *http.Request) {
}

func DeleteReminder(w http.ResponseWriter, req *http.Request) {
}

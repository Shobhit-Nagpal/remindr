package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Shobhit-Nagpal/remindr/internal/db"
	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
)

var managerKey string = "manager"
var dbKey string = "db"

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func respondWithError(w http.ResponseWriter, code int, msg string) error {
	return respondWithJSON(w, code, map[string]string{"error": msg})
}


func getManager(r *http.Request) *jobs.JobManager {
	manager, ok := r.Context().Value(managerKey).(*jobs.JobManager)
	if !ok {
		return nil
	}

	return manager
}

func getDB(r *http.Request) *db.DB {
	manager, ok := r.Context().Value(dbKey).(*db.DB)
	if !ok {
		return nil
	}

	return manager
}

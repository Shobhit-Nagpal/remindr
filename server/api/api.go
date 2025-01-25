package api

import (
	"net/http"

	"github.com/Shobhit-Nagpal/remindr/internal/db"
	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
	"github.com/Shobhit-Nagpal/remindr/server/handler"
	"github.com/Shobhit-Nagpal/remindr/server/middleware"
)

func NewServer(database *db.DB, manager *jobs.JobManager) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.Index)
	mux.HandleFunc("GET /api/healthz", handler.Health)

	mux.HandleFunc("GET /api/reminders", handler.GetReminders)
	mux.HandleFunc("POST /api/reminders", handler.CreateReminder)
	mux.HandleFunc("DELETE /api/reminders", handler.DeleteReminder)

	handler := middleware.DB(mux, database)
	handler = middleware.Manager(handler, manager)
	handler = middleware.Logger(handler)

	server := &http.Server{
		Handler: handler,
		Addr:    "localhost:5678",
	}

	return server
}

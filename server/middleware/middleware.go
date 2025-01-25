package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Shobhit-Nagpal/remindr/internal/jobs"
)

func Manager(next http.Handler, manager *jobs.JobManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		ctx = context.WithValue(ctx, "manager", manager)

		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("\nMethod: %s\nPath: %s\nDuration: %s\n\n", req.Method, req.URL.EscapedPath(), time.Since(start))
	})
}


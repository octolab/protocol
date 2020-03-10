package middleware

import (
	"context"
	"net/http"
	"time"

	"go.octolab.org/toolkit/protocol/http/header"
)

// Deadline returns the Middleware to inject a deadline timestamp
// into the request context.
func Deadline(fallback time.Duration, corrector func(time.Time, bool) time.Time) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			deadline := corrector(header.Deadline(req.Header, fallback))
			ctx, cancel := context.WithDeadline(req.Context(), deadline)
			handler.ServeHTTP(rw, req.WithContext(ctx))
			cancel()
		})
	}
}

// Timeout returns the Middleware to inject a timeout duration
// into the request context.
func Timeout(fallback time.Duration, corrector func(time.Duration, bool) time.Duration) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			timeout := corrector(header.Timeout(req.Header, fallback))
			ctx, cancel := context.WithTimeout(req.Context(), timeout)
			handler.ServeHTTP(rw, req.WithContext(ctx))
			cancel()
		})
	}
}
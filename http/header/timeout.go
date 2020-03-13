package header

import (
	"net/http"
	"time"
)

const (
	XDeadline = "X-Deadline"
	XTimeout  = "X-Timeout"
)

// Deadline returns the deadline timestamp from the header.
// It tries to use the timeout duration from the header
// if the deadline is invalid or doesn't exist in the header.
// Last it tries to use the fallback duration to calculate
// the appropriate deadline.
func Deadline(header http.Header, fallback time.Duration) (time.Time, bool) {
	deadline, err := time.Parse(time.RFC3339Nano, header.Get(XDeadline))
	if err == nil {
		return deadline, true
	}
	if duration, present := Timeout(header, fallback); present {
		return time.Now().Add(duration), true
	}
	return time.Now().Add(fallback), false
}

// Timeout returns the timeout duration from the header.
// It tries to use the fallback duration if the timeout is invalid
// or doesn't exist in the header.
func Timeout(header http.Header, fallback time.Duration) (time.Duration, bool) {
	duration, err := time.ParseDuration(header.Get(XTimeout))
	if err != nil {
		return fallback, false
	}
	return duration, true
}

package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	CacheControlHeader = "Cache-Control"
	ContentTypeHeader  = "Content-Type"
	XDeadlineHeader    = "X-Deadline"
	XStrictHeader      = "X-Strict"
	XTimeoutHeader     = "X-Timeout"
)

// Header extends built-in http.Header.
type Header http.Header

// Deadline returns the deadline duration from the header or the present duration.
// It also tries to use the timeout duration from the header if the deadline is invalid.
func (header Header) Deadline(fallback time.Duration) (time.Time, bool) {
	deadline, err := time.Parse(time.RFC3339Nano, http.Header(header).Get(XDeadlineHeader))
	if err == nil {
		return deadline, true
	}
	if duration, present := header.Timeout(fallback); present {
		return time.Now().Add(duration), true
	}
	return time.Now().Add(fallback), false
}

// NoCache returns true if the header has no-cache duration of cache control.
func (header Header) NoCache() bool {
	return strings.EqualFold(http.Header(header).Get(CacheControlHeader), "no-cache")
}

// Strict returns true if the header has this duration.
func (header Header) Strict() bool {
	var strict bool
	if v := http.Header(header).Get(XStrictHeader); v != "" {
		strict, _ = strconv.ParseBool(v)
	}
	return strict
}

// Timeout returns the timeout duration from the header or the present duration.
func (header Header) Timeout(fallback time.Duration) (time.Duration, bool) {
	duration, err := time.ParseDuration(http.Header(header).Get(XTimeoutHeader))
	if err != nil {
		return fallback, false
	}
	return duration, true
}

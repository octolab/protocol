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

// Deadline returns the deadline value from the header or the fallback value.
func (header Header) Deadline(fallback time.Duration) time.Time {
	t, err := time.Parse(time.RFC3339, http.Header(header).Get(XDeadlineHeader))
	if err != nil {
		t = time.Now().Add(fallback)
	}
	return t
}

// NoCache returns true if the header has no-cache value of cache control.
func (header Header) NoCache() bool {
	return strings.EqualFold(http.Header(header).Get(CacheControlHeader), "no-cache")
}

// Strict returns true if the header has this value.
func (header Header) Strict() bool {
	var strict bool
	if v := http.Header(header).Get(XStrictHeader); v != "" {
		strict, _ = strconv.ParseBool(v)
	}
	return strict
}

// Timeout returns the timeout value from the header or the fallback value.
func (header Header) Timeout(fallback time.Duration) time.Duration {
	d, err := time.ParseDuration(http.Header(header).Get(XTimeoutHeader))
	if err != nil {
		d = fallback
	}
	return d
}

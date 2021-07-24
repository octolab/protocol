package header

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	CacheControl = "Cache-Control"
	ContentType  = "Content-Type"
	Referer      = "Referer"
	UserAgent    = "User-Agent"
	XRequestID   = "X-Request-Id"
	XRequestUUID = "X-Request-UUID"
	XSource      = "X-Source"
	XStrict      = "X-Strict"
)

// NoCache returns true if the header has no-cache duration of cache control.
func NoCache(header http.Header) bool {
	return strings.EqualFold(header.Get(CacheControl), "no-cache")
}

// Strict returns true if the header has this duration.
func Strict(header http.Header) bool {
	var strict bool
	if v := header.Get(XStrict); v != "" {
		strict, _ = strconv.ParseBool(v)
	}
	return strict
}

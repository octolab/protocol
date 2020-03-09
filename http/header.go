package http

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	CacheControlHeader = "Cache-Control"
	ContentTypeHeader  = "Content-Type"
	XStrictHeader      = "X-Strict"
)

// Header extends built-in http.Header.
type Header http.Header

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

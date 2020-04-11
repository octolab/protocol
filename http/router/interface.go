package router

import "net/http"

// Interface is an HTTP request multiplexer.
// It is compatible with net/http.ServerMux.
type Interface interface {
	http.Handler
	// Handle registers the handler for the given pattern.
	Handle(string, http.Handler)
	// HandleFunc registers the handler function for the given pattern.
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}

package rest

import "net/http"

// Interface is an HTTP request multiplexer.
// It is compatible with github.com/go-chi/chi.Mux.
type Interface interface {
	http.Handler
	// Handle registers the handler for the given pattern.
	Handle(string, http.Handler)
	// HandleFunc registers the handler function for the given pattern.
	HandleFunc(string, http.HandlerFunc)

	// Connect adds the route that matches a CONNECT HTTP method to
	// execute the handler.
	Connect(string, http.HandlerFunc)
	// Delete adds the route that matches a DELETE HTTP method to
	// execute the handler.
	Delete(string, http.HandlerFunc)
	// Get adds the route that matches a GET HTTP method to
	// execute the handler.
	Get(string, http.HandlerFunc)
	// Head adds the route that matches a HEAD HTTP method to
	// execute the handler.
	Head(string, http.HandlerFunc)
	// Options adds the route that matches a OPTIONS HTTP method to
	// execute the handler.
	Options(string, http.HandlerFunc)
	// Patch adds the route that matches a PATCH HTTP method to
	// execute the handler.
	Patch(string, http.HandlerFunc)
	// Post adds the route that matches a POST HTTP method to
	// execute the handler.
	Post(string, http.HandlerFunc)
	// Put adds the route that matches a PUT HTTP method to
	// execute the handler.
	Put(string, http.HandlerFunc)
	// Trace adds the route that matches a TRACE HTTP method to
	// execute the handler.
	Trace(string, http.HandlerFunc)
}

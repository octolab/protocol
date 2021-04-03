package router

import "net/http"

// Handler is a http handler with a specified path.
// It is compatible with net/http.ServeMux.
type Handler func() (path string, handler http.Handler)

// HandlerFunc is a http handler with a specified path.
// It is compatible with net/http.ServeMux.
type HandlerFunc func() (path string, handler http.HandlerFunc)

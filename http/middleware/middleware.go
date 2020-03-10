package middleware

import "net/http"

// Middleware defines the Handler wrapper.
type Middleware func(http.Handler) http.Handler

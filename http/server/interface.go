package server

import (
	"context"
	"io"
	"net"
)

// Interface represents a generic server to listen a network protocol.
// It is compatible with net/http.Server.
type Interface interface {
	io.Closer
	// Serve accepts incoming connections on the Listener and serves them.
	Serve(net.Listener) error
	// ListenAndServe listens a network protocol and serves it.
	ListenAndServe() error
	// RegisterOnShutdown registers a function to call on Shutdown.
	RegisterOnShutdown(func())
	// Shutdown tries to do a graceful shutdown.
	Shutdown(context.Context) error
}

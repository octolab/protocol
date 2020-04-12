package server

import (
	"context"

	"github.com/pkg/errors"
)

// Run runs ListenAndServe in separated goroutine and listens shutdown signal.
// It returns ListenAndServe' error or Shutdown' error if signal is received.
func Run(server Interface, shutdown chan context.Context) error {
	result := make(chan error, 1)

	go func() {
		var err error

		defer func() {
			result <- err
			close(result)
		}()
		defer func() {
			if r := recover(); r != nil {
				err = errors.Errorf("panic unexpected: %#+v", r)
			}
		}()

		err = errors.Wrap(server.ListenAndServe(), "server: listen and serve")
	}()

	select {
	case err := <-result:
		return err
	case ctx := <-shutdown:
		return errors.Wrap(server.Shutdown(ctx), "server: shutdown")
	}
}

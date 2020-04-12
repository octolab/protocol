package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/http/server"
)

//go:generate mockgen -package $GOPACKAGE -destination mock_server_test.go go.octolab.org/toolkit/protocol/http/server Interface

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := map[string]struct {
		server   func() Interface
		shutdown func() chan context.Context
		assert   func(assert.TestingT, error, ...interface{}) bool
	}{
		"listen and serve error": {
			func() Interface {
				mock := NewMockInterface(ctrl)
				mock.EXPECT().
					ListenAndServe().
					Return(errors.New("listen and serve"))
				return mock
			},
			func() chan context.Context { return nil },
			assert.Error,
		},
		"listen and serve panic": {
			func() Interface {
				mock := NewMockInterface(ctrl)
				mock.EXPECT().
					ListenAndServe().
					Do(func() { panic("bad server") })
				return mock
			},
			func() chan context.Context { return nil },
			assert.Error,
		},
		"shutdown error": {
			func() Interface {
				timer := time.NewTimer(time.Hour)
				mock := NewMockInterface(ctrl)
				mock.EXPECT().
					ListenAndServe().
					Do(func() { <-timer.C }).
					Return(errors.New("shutdown"))
				mock.EXPECT().
					Shutdown(context.Background()).
					Do(func(context.Context) { _ = timer.Stop() }).
					Return(errors.New("shutdown"))
				return mock
			},
			func() chan context.Context {
				ch := make(chan context.Context, 1)
				ch <- context.Background()
				return ch
			},
			assert.Error,
		},
		"graceful shutdown": {
			func() Interface {
				timer := time.NewTimer(time.Hour)
				mock := NewMockInterface(ctrl)
				mock.EXPECT().
					ListenAndServe().
					Do(func() { <-timer.C }).
					Return(errors.New("shutdown"))
				mock.EXPECT().
					Shutdown(context.Background()).
					Do(func(context.Context) { _ = timer.Stop() }).
					Return(nil)
				return mock
			},
			func() chan context.Context {
				ch := make(chan context.Context, 1)
				ch <- context.Background()
				return ch
			},
			assert.NoError,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.assert(t, Run(test.server(), test.shutdown()))
		})
	}
}

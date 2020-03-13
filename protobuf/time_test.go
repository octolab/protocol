package protobuf_test

import (
	"math"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/protobuf"
)

func TestTime(t *testing.T) {
	tests := map[string]struct {
		timestamp *timestamp.Timestamp
		assert    func(assert.TestingT, *timestamp.Timestamp)
	}{
		"nil pointer": {
			timestamp: nil,
			assert: func(t assert.TestingT, tsp *timestamp.Timestamp) {
				assert.Nil(t, Time(tsp))
			},
		},
		"normal use": {
			timestamp: new(timestamp.Timestamp),
			assert: func(t assert.TestingT, tsp *timestamp.Timestamp) {
				assert.NotNil(t, Time(tsp))
			},
		},
		"invalid timestamp": {
			timestamp: &timestamp.Timestamp{Seconds: -1, Nanos: -1},
			assert: func(t assert.TestingT, ts *timestamp.Timestamp) {
				assert.Panics(t, func() { Time(ts) })
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.assert(t, test.timestamp)
		})
	}
}

func TestTimestamp(t *testing.T) {
	tests := map[string]struct {
		time   *time.Time
		assert func(assert.TestingT, *time.Time)
	}{
		"nil pointer": {
			time: nil,
			assert: func(t assert.TestingT, tp *time.Time) {
				assert.Nil(t, Timestamp(tp))
			},
		},
		"normal use": {
			time: new(time.Time),
			assert: func(t assert.TestingT, tp *time.Time) {
				assert.NotNil(t, Timestamp(tp))
			},
		},
		"invalid time": {
			time: func() *time.Time {
				tp := time.Now().AddDate(-math.MaxInt32, -math.MaxInt32, -math.MaxInt32)
				return &tp
			}(),
			assert: func(t assert.TestingT, tp *time.Time) {
				assert.Panics(t, func() { Timestamp(tp) })
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.assert(t, test.time)
		})
	}
}

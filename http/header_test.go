package http_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/protocol/http"
)

func TestHeader_NoCache(t *testing.T) {
	tests := map[string]struct {
		header   Header
		expected bool
	}{
		"exists in header": {
			Header{CacheControlHeader: []string{"no-cache"}},
			true,
		},
		"exists in header, case insensitive": {
			Header{CacheControlHeader: []string{"No-Cache"}},
			true,
		},
		"empty duration": {
			nil,
			false,
		},
		"another duration": {
			Header{CacheControlHeader: []string{"only-if-cached"}},
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.header.NoCache())
		})
	}
}

func TestHeader_Strict(t *testing.T) {
	tests := map[string]struct {
		header   Header
		expected bool
	}{
		"exists in header, string": {
			Header{XStrictHeader: []string{"true"}},
			true,
		},
		"exists in header, numeric": {
			Header{XStrictHeader: []string{"1"}},
			true,
		},
		"exists in header, case insensitive": {
			Header{XStrictHeader: []string{"True"}},
			true,
		},
		"empty duration": {
			nil,
			false,
		},
		"invalid duration": {
			Header{XStrictHeader: []string{"invalid"}},
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.header.Strict())
		})
	}
}

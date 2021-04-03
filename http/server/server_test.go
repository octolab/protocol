package server_test

import (
	"net/http"

	. "go.octolab.org/toolkit/protocol/http/server"
)

var _ Interface = &http.Server{}

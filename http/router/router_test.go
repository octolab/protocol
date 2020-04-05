package router_test

import (
	"net/http"

	. "go.octolab.org/toolkit/protocol/http/router"
)

var _ Interface = http.NewServeMux()

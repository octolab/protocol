package rest_test

import (
	"github.com/go-chi/chi"

	. "go.octolab.org/toolkit/protocol/http/router/rest"
)

var _ Interface = &chi.Mux{}

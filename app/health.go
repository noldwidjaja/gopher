package app

import "net/http"

type health struct {
	Message string `json:"message"`
}

func healthcheck() http.HandlerFunc {
	r := newResponse(
		health{"great"},
		"ok",
		200,
	)
	return r.HandleJSONResponse()
}

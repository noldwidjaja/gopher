package app

import "net/http"

type health struct {
	Message string `json:"message"`
}

// HealthCheck : used to test routing
func HealthCheck() http.HandlerFunc {
	r := NewResponse(
		health{"great"},
		"ok",
		200,
	)
	return r.HandleJSONResponse()
}

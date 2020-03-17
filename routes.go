package main

import (
	"github.com/noldwidjaja/money-manager/app"
)

// Routes is where routing will be inserted
func (s *Server) routes() {
	s.Get("/health", app.HealthCheck())
}

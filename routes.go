package main

// Routes is where routing will be inserted
func (s *Server) routes() {
	s.Get("/health", app.healthcheck())
}

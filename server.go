package main

import (
	"fmt"
	"net/http"
)

// Router serves http
type Router struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

// Server creates golang server
type Server struct {
	Router *Router
	port   string
}

// NewServer constructs the server with the router inside the server
func NewServer(portNumber string) *Server {
	s := &Server{
		Router: NewRouter(),
		port:   ":" + portNumber,
	}
	return s
}

// NewRouter creates instance of Router
func NewRouter() *Router {
	router := new(Router)
	router.handlers = make(map[string]func(http.ResponseWriter, *http.Request))
	return router
}

// ServeHTTP is called for every connection
func (s *Router) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	f, ok := s.handlers[key(r.Method, r.URL.Path)]
	if !ok {
		bad(w)
		return
	}
	f(w, r)
}

// Get sets get handler
func (s *Server) Get(path string, f http.HandlerFunc) {
	s.Router.handlers[key("GET", path)] = f
}

// Post sets post handler
func (s *Server) Post(path string, f http.HandlerFunc) {
	s.Router.handlers[key("POST", path)] = f
}

// Delete sets delete handler
func (s *Server) Delete(path string, f http.HandlerFunc) {
	s.Router.handlers[key("DELETE", path)] = f
}

// Put sets put handler
func (s *Server) Put(path string, f http.HandlerFunc) {
	s.Router.handlers[key("PUT", path)] = f
}

func key(method, path string) string {
	return fmt.Sprintf("%s:%s", method, path)
}

func bad(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error":"not found"}`))
}

// Serve serves the server for development
func (s *Server) serve() error {
	return http.ListenAndServe(s.port, s.Router)
}

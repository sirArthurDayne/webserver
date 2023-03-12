package main

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

//generar un servidor mediante puerto y router
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

//combinar ruta con Handler
func (s *Server) Handle(path string, handler http.HandlerFunc) {
    s.router.rules[path] = handler
}

//escuchar las peticiones
func (s *Server) Listen() error {
    http.Handle("/", s.router)
	error := http.ListenAndServe(s.port, nil)
	if error != nil {
		return error
	}
	return nil
}


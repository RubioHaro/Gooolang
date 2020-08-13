package main

import (
	"fmt"
	"net/http"
)

// Server This is a struc from type server /(parameters: port, router)
type Server struct {
	port   string
	router *Router
}

func newServer(port string) *Server {
	fmt.Println("Iniciando servidor en el puerto:", port)
	return &Server{
		port:   port,
		router: newRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, existe := s.router.rules[path]
	if !existe {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) addMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (s *Server) listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}

	return nil
}

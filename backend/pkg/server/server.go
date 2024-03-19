package server

import "net/http"


type Server struct {
	Addr    string
	handler http.Handler
}


func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}


func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.Addr, s.handler)
}

package api

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

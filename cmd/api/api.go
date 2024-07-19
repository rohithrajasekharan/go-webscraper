package api

import (
	"log"
	"net/http"

	"github.com/rohithrajasekharan/web-scraper/service"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()
	serviceHandler := service.NewHandler()
	serviceHandler.RegisterRoutes(mux)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

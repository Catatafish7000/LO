package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	cache      Getter
	httpServer *http.Server
}

func New(port string, db Getter) *Server {
	serv := Server{
		cache: db,
		httpServer: &http.Server{
			Addr: port,
		},
	}
	return &serv
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")
	ord, err := s.cache.GetInformation(id)
	if err != nil {
		log.Print("error creating json")
	} else {
		fmt.Fprintf(w, string(ord))
	}
}

func (s *Server) Serve() error {
	http.HandleFunc("/", s.handler)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

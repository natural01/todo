package pkg

import (
	"net/http"
)

type Server struct {
	HttpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	s.HttpServer = &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	return s.HttpServer.ListenAndServe()
}

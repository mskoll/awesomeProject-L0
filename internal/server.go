package internal

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

// Run запуск сервера
func (s *Server) Run(port string, handler http.Handler) error {

	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown остановка сервера
func (s *Server) Shutdown(contx context.Context) error {

	return s.httpServer.Shutdown(contx)
}

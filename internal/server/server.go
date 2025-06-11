package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func New(logger *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.HomePageHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger:     logger,
		httpServer: httpServer,
	}

}

func (s *Server) Start() error {
	s.logger.Printf("Старт сервера на порт %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

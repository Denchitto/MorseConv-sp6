package server

import (
	"MorseConv-sp6/internal/handlers"
	"log"
	"net/http"
	"time"
)

type CustomServer struct {
	Logger *log.Logger
	Server *http.Server
}

func MakeRouter(lg *log.Logger) (s *CustomServer) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	var serv = &CustomServer{
		Logger: lg,
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     lg,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	return serv
}

package http

import (
	"log"
	"net/http"

	"github.com/Diverstt/Yandex_Sprint1/internal/application"
	"github.com/Diverstt/Yandex_Sprint1/internal/infrastructure/config"
)

// StartServer запускает сервер на определенном порте
func StartServer(handler application.CalcHandler) {
	port := config.GetPort()

	mux := http.NewServeMux()

	SetupRoutes(mux, handler)

	log.Printf("Starting server on %s", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}

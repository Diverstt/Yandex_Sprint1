package http

import (
	"net/http"

	"github.com/Diverstt/Yandex_Sprint1/internal/application"
)

// SetupRoutes регистрирует все маршруты для HTTP-сервера.
func SetupRoutes(mux *http.ServeMux, handler application.CalcHandler) {
	mainHandler := http.HandlerFunc(handler.HandlerCalc)
	mux.HandleFunc("/calculate", application.MethodMiddleware(mainHandler))

	mux.Handle("/", http.FileServer(http.Dir("./static")))
}

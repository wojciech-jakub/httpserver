package routes

import (
	"httpserver/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// API endpoints
	mux.Handle("/random/mean", http.HandlerFunc(handlers.ApiHandler))

	return mux
}

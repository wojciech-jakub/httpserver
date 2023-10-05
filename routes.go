package main

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// API endpoints
	mux.Handle("/random/mean", http.HandlerFunc(apiHandler))

	return mux
}

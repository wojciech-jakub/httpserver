package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"httpserver/config"
	"httpserver/routes"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	routesHandler := routes.SetupRoutes()
	defaultServerConfig := config.DefaultServerConfig()

	server := &http.Server{
		Addr:         defaultServerConfig.Addr,
		Handler:      routesHandler,
		ReadTimeout:  defaultServerConfig.ReadTimeout,
		WriteTimeout: defaultServerConfig.WriteTimeout,
		IdleTimeout:  defaultServerConfig.IdleTimeout,
		BaseContext: func(listener net.Listener) context.Context {
			ctx = context.WithValue(ctx, defaultServerConfig.Addr, listener.Addr().String())
			return ctx
		},
	}

	go func() {
		fmt.Printf("Starting server on %s\n", defaultServerConfig.Addr)
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case <-c:
		fmt.Println("Shutting down server...")

		// Create a context with a deadline for graceful shutdown
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		// Shutdown the server
		if err := server.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("Server shutdown error: %v\n", err)
		}
	}
}

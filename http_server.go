package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

func startHttpServer(ctx context.Context) func() {
	routesHandler := SetupRoutes()
	defaultServerConfig := DefaultServerConfig()

	httpServer := &http.Server{
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
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	return func() {
		fmt.Printf("Shuting down server")
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("Server shutdown error: %v\n", err)
		}
	}
}

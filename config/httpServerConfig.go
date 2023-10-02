package config

import (
	"fmt"
	"time"
)

// ServerConfig represents the configuration for the HTTP server.
type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Addr:         fmt.Sprintf("%s:%s", "127.0.0.1", "8080"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

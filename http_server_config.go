package main

import (
	"fmt"
	"github.com/Netflix/go-env"
	"time"
)

type Environment struct {
	Address            string        `env:"HTTP_SERVER_ADDRESS"`
	Port               string        `env:"HTTP_SERVER_PORT"`
	ServerReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT"`
	ServerIdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT"`

	Extras env.EnvSet
}

// ServerConfig represents the configuration for the HTTP server.
type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Addr:         fmt.Sprintf("%s:%s", "0.0.0.0", "8080"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

//https://github.com/Netflix/go-env

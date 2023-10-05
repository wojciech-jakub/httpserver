package main

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	stopHttpServer := startHttpServer(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	stopHttpServer()
}

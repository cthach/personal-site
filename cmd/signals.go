package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handleSignals(ctx context.Context) error {
	ch := make(chan os.Signal, 2)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		return ctx.Err()

	case sig := <-ch:
		fmt.Printf("Received: %s\n", sig)
		return fmt.Errorf("received signal: %s", sig)
	}
}

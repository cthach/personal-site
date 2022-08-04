package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/magefile/mage/target"
	"golang.org/x/sync/errgroup"
)

// Starts the development server and will restart the server on any changes
func Start() error {
	for {
		eg, ctx := errgroup.WithContext(context.Background())

		eg.Go(func() error {
			return detectChangesInCurrentDirectory(ctx)
		})

		eg.Go(func() error {
			// FIXME: Child commands still living after killed
			// TODO: Open browser after starting
			return run(ctx)
		})

		if err := eg.Wait(); err != nil {
			if err.Error() == "reload" {
				fmt.Println("File changes detected, reloading")
				continue
			}

			return fmt.Errorf("exited: %w", err)
		}
	}
}

func run(ctx context.Context) error {
	cmd := exec.CommandContext(
		ctx,
		"go",
		"run",
		"./cmd",
		"-addr", ":0",
	)

	// Pipe all output to stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func detectChangesInCurrentDirectory(ctx context.Context) error {
	now := time.Now()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-time.After(2 * time.Second):
			newChanges, err := target.DirNewer(
				now,
				"cmd",
				"http",
			)
			if err != nil {
				return fmt.Errorf("watch for changes: %w", err)
			}
			if newChanges {
				return fmt.Errorf("reload")
			}
		}
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/cthach/personal-site/http"
)

var (
	code int

	addr string
)

func main() {
	defer os.Exit(code)

	flag.StringVar(&addr, "addr", ":8080", "address to listen on")
	flag.Parse()

	lis, err := net.Listen("tcp4", addr)
	if err != nil {
		fmt.Printf("listen: %s\n", err)
		code = 1
		return
	}
	defer lis.Close()

	svr := &http.Server{Listener: lis}

	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		return handleSignals(ctx)
	})

	eg.Go(svr.ListenAndServe)

	eg.Go(func() error {
		<-ctx.Done()
		defer fmt.Println("shutdown gracefully")
		return svr.GracefulShutdown(context.TODO())
	})

	fmt.Printf("http server listening on http://%s\n", lis.Addr().String())

	if err := eg.Wait(); err != nil {
		fmt.Printf("server stopped: %s\n", err)

		if !strings.Contains(err.Error(), "received signal") {
			code = 1
			return
		}
	}
}

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

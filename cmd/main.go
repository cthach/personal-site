package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/cthach/personal-site/http"
)

func main() {
	svr := &http.Server{Addr: ":8080"}

	eg, _ := errgroup.WithContext(context.Background())

	eg.Go(svr.ListenAndServe)

	// TODO: Handle graceful shutdown

	if err := eg.Wait(); err != nil {
		fmt.Printf("server stopped: %s", err)
	}
}

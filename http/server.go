package http

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"
)

//go:embed static
var content embed.FS

// Server wraps http.Server.
type Server struct {
	Listener net.Listener
	http     *http.Server
}

// ListenAndServe will listen and serve on the server address. Blocks until the server is stopped.
func (s *Server) ListenAndServe() error {
	static, err := fs.Sub(content, "static")
	if err != nil {
		return fmt.Errorf("get embedded static files subtree: %w", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.FS(static)))

	s.http = &http.Server{
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	return s.http.Serve(s.Listener)
}

// GracefulShutdown will gracefully shutdown the server.
func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

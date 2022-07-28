package http

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//go:embed static
var content embed.FS

// Server wraps http.Server.
type Server struct {
	Addr string
	http *http.Server
}

// ListenAndServe will listen and serve on the server address. Blocks until the server is stopped.
func (s *Server) ListenAndServe() error {
	static, err := fs.Sub(content, "static")
	if err != nil {
		return fmt.Errorf("get embedded static files subtree: %w", err)
	}

	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.FileServer(http.FS(static)))

	s.http = &http.Server{
		Addr:           s.Addr,
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.http.ListenAndServe()
}

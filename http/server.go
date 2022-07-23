package http

import (
	_ "embed"
	"net/http"
	"time"
)

//go:embed static/index.html
var content []byte

type Server struct {
	Addr string
	http *http.Server
}

func (s *Server) ListenAndServe() error {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(content)
		},
	)

	s.http = &http.Server{
		Addr:           s.Addr,
		Handler:        nil, // Use DefaultServeMux
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.http.ListenAndServe()
}

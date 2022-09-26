package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/server"
)

const DefaultShutdownTimeout = time.Minute

var _ server.Server = (*Server)(nil)

// Server gin implement http server
type Server struct {
	ShutdownTimeout time.Duration
	srv             *http.Server
}

type Options func(*Server)

// NewServer creates a new server instance with default settings
func NewServer(e *gin.Engine, addr string, options ...Options) *Server {
	ser := Server{
		ShutdownTimeout: DefaultShutdownTimeout,
		srv: &http.Server{
			Addr:    addr,
			Handler: e,
		},
	}

	for _, option := range options {
		option(&ser)
	}

	return &ser
}

// WithGracefulShutdownDuration duration of graceful shutdown.
//
// Deprecated: this function is deprecated, because it's name to long, replace it with `WithShutdownTimeout`
func WithGracefulShutdownDuration(gracefulShutdownDuration time.Duration) Options {
	return WithShutdownTimeout(gracefulShutdownDuration)
}

// WithShutdownTimeout duration of graceful shutdown
func WithShutdownTimeout(duration time.Duration) Options {
	return func(server *Server) {
		server.ShutdownTimeout = duration
	}
}

// Start to start the server and wait for it to listen on the given address
func (h *Server) Start() (err error) {
	return h.srv.ListenAndServe()
}

// Stop stops the server and close with graceful shutdown duration
//
// Deprecated: use shutdown instead.
func (h *Server) Stop() error {
	return h.Shutdown()
}

// Shutdown shuts down the server and close with graceful shutdown duration
func (h *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), h.ShutdownTimeout)
	defer cancel()

	return h.srv.Shutdown(ctx)
}

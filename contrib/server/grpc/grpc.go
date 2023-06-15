package grpc

import (
	"net"

	"github.com/segmentfault/pacman/server"
	"google.golang.org/grpc"
)

var _ server.Server = (*Server)(nil)

// Server grpc server
type Server struct {
	srv  *grpc.Server
	addr string
}

type Options func(*Server)

// NewServer creates a new server instance with default settings
func NewServer(grpcServer *grpc.Server, addr string, options ...Options) *Server {
	ser := Server{
		srv:  grpcServer,
		addr: addr,
	}

	for _, option := range options {
		option(&ser)
	}
	return &ser
}

// Start to start the server and listen on the given address
func (h *Server) Start() (err error) {
	lis, err := net.Listen("tcp", h.addr)
	if err != nil {
		return err
	}
	if err = h.srv.Serve(lis); err != nil {
		return err
	}
	return nil
}

// Shutdown shuts down the server
func (h *Server) Shutdown() error {
	h.srv.GracefulStop()
	return nil
}

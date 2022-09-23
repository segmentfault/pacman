package server

// Server is transport server.
type Server interface {
	Start() error
	Stop() error
}

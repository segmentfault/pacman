package pacman

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentfault/pacman/log"
	"github.com/segmentfault/pacman/server"
)

// Application is the main struct of the application
type Application struct {
	id      string
	name    string
	version string
	servers []server.Server
	signals []os.Signal
}

// Option application support option
type Option func(application *Application)

// NewApp creates a new Application
func NewApp(ops ...Option) *Application {
	app := &Application{}
	for _, op := range ops {
		op(app)
	}
	// default random id
	if len(app.id) == 0 {
		bytes := make([]byte, 24)
		_, _ = rand.Read(bytes)
		app.id = hex.EncodeToString(bytes)
	}
	// default accept signals
	if len(app.signals) == 0 {
		app.signals = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	}
	return app
}

// WithID application add id
func WithID(id string) func(application *Application) {
	return func(application *Application) {
		application.id = id
	}
}

// WithName application add name
func WithName(name string) func(application *Application) {
	return func(application *Application) {
		application.name = name
	}
}

// WithVersion application add version
func WithVersion(version string) func(application *Application) {
	return func(application *Application) {
		application.version = version
	}
}

// WithServer application add server
func WithServer(servers ...server.Server) func(application *Application) {
	return func(application *Application) {
		application.servers = servers
	}
}

// WithSignals application add listen signals
func WithSignals(signals []os.Signal) func(application *Application) {
	return func(application *Application) {
		application.signals = signals
	}
}

// Run application run
func (app *Application) Run() error {
	if len(app.servers) == 0 {
		return nil
	}

	for _, s := range app.servers {
		go func(srv server.Server) {
			if err := srv.Start(); err != nil {
				log.Errorf("failed to start server, err: %s", err)
			}
		}(s)
	}

	quit := make(chan os.Signal, 8)
	signal.Notify(quit, app.signals...)
	<-quit
	return nil
}

// Stop application stop
func (app *Application) Stop() error {
	for _, s := range app.servers {
		go func(ser server.Server) {
			if err := ser.Stop(); err != nil {
				log.Errorf("failed to stop server, err: %s", err)
			}
		}(s)
	}
	return nil
}

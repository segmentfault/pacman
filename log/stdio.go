package log

import (
	"io"
	"log"
)

var _ Logger = (*stdLogger)(nil)

type stdLogger struct {
	log *log.Logger
}

func (s *stdLogger) Debug(v ...any) {
	s.log.Println(v)
}

func (s *stdLogger) Debugf(format string, v ...any) {
	s.log.Printf(format, v)
}

func (s *stdLogger) Info(v ...any) {
	s.log.Println(v)
}

func (s *stdLogger) Infof(format string, v ...any) {
	s.log.Printf(format, v)
}

func (s *stdLogger) Warn(v ...any) {
	s.log.Println(v)
}

func (s *stdLogger) Warnf(format string, v ...any) {
	s.log.Printf(format, v)
}

func (s *stdLogger) Error(v ...any) {
	s.log.Println(v)
}

func (s *stdLogger) Errorf(format string, v ...any) {
	s.log.Printf(format, v)
}

// NewStdLogger new a logger with writer.
func NewStdLogger(w io.Writer) Logger {
	return &stdLogger{
		log: log.New(w, "", 0),
	}
}

package log

import (
	"log"
)

// DefaultLogger is default logger.
var DefaultLogger = NewStdLogger(log.Writer())

// Logger logger interface
type Logger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
}

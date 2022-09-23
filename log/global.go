package log

// global
var global Logger

func init() {
	global = DefaultLogger
}

// SetLogger set logger
func SetLogger(log Logger) {
	global = log
}

// GetLogger get logger
func GetLogger() Logger {
	return global
}

// Debug log
func Debug(v ...any) {
	global.Debug(v...)
}

// Debugf log
func Debugf(format string, v ...any) {
	global.Debugf(format, v...)
}

// Info log
func Info(v ...any) {
	global.Info(v...)
}

// Infof log
func Infof(format string, v ...any) {
	global.Infof(format, v...)
}

// Warn log
func Warn(v ...any) {
	global.Warn(v...)
}

// Warnf log
func Warnf(format string, v ...any) {
	global.Warnf(format, v...)
}

// Error log
func Error(v ...any) {
	global.Error(v...)
}

// Errorf log
func Errorf(format string, v ...any) {
	global.Errorf(format, v...)
}

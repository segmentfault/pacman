package zap

import "time"

type LogOption func(*Logger)

// WithName set file name
func WithName(name string) LogOption {
	return func(l *Logger) {
		l.conf.name = name
	}
}

// WithPath set file path
func WithPath(path string) LogOption {
	return func(l *Logger) {
		l.conf.path = path
	}
}

// WithMaxAge set max age
func WithMaxAge(maxAge time.Duration) LogOption {
	return func(l *Logger) {
		l.conf.maxAge = maxAge
	}
}

// WithRotationTime set rotation time
func WithRotationTime(rotationTime time.Duration) LogOption {
	return func(l *Logger) {
		l.conf.rotationTime = rotationTime
	}
}

// WithoutStd no output stdout
func WithoutStd() LogOption {
	return func(l *Logger) {
		l.conf.stdout = false
	}
}

// WithCallerFullPath log caller full path
func WithCallerFullPath() LogOption {
	return func(l *Logger) {
		l.conf.callerFullPath = true
	}
}

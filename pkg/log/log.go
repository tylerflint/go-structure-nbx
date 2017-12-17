// Package log provides a simple logging interface with the means to log
// to an 'io.Writer.'
package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Logger defines a simple logger object.
type Logger struct {
	// The logs are `fmt.Fprintf`d to this. It's common to set this to a
	// file, or leave it default which is `os.Stderr`.
	Out io.Writer
	// The logging level the logger should log at. This is typically (and defaults
	// to) `Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged.
	Level level
	// The date format which the logger should print the timestamp as. If not set,
	// timestamp is not printed.
	Date string
}

type level uint32

const (
	// FatalLevel level. Logs. Used for critical errors.
	FatalLevel = iota
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Usually only enabled when debugging. Very very verbose logging.
	TraceLevel
)

// New creates a new logger object with sane defaults.
func New() *Logger {
	return &Logger{
		Out:   os.Stderr,
		Level: InfoLevel,
		Date:  "2006-01-02 15:04:05.00000",
	}
}

// Printf writes a the log the logger's Out.
func (logger Logger) Printf(format string, args ...interface{}) {
	if format[len(format)-1] != '\n' {
		format += "\n"
	}
	fmt.Fprintf(logger.Out, format, args...)
}

// Trace writes a log to the loggers 'Out' if its 'Level' is at least trace.
func (logger Logger) Trace(format string, args ...interface{}) {
	if logger.Level >= TraceLevel {
		format = fmt.Sprintf("%5s  %s", "TRACE", format)
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
		if logger.Date != "" {
			format = fmt.Sprintf("%s %s", time.Now().UTC().Format(logger.Date), format)
		}
		fmt.Fprintf(logger.Out, format, args...)
	}
}

// Debug writes a log to the loggers 'Out' if its 'Level' is at least debug.
func (logger Logger) Debug(format string, args ...interface{}) {
	if logger.Level >= DebugLevel {
		format = fmt.Sprintf("%5s  %s", "DEBUG", format)
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
		if logger.Date != "" {
			format = fmt.Sprintf("%s %s", time.Now().UTC().Format(logger.Date), format)
		}
		fmt.Fprintf(logger.Out, format, args...)
	}
}

// Info writes a log to the loggers 'Out' if its 'Level' is at least info.
func (logger Logger) Info(format string, args ...interface{}) {
	if logger.Level >= InfoLevel {
		format = fmt.Sprintf("%5s  %s", "INFO", format)
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
		if logger.Date != "" {
			format = fmt.Sprintf("%s %s", time.Now().UTC().Format(logger.Date), format)
		}
		fmt.Fprintf(logger.Out, format, args...)
	}
}

// Warn writes a log to the loggers 'Out' if its 'Level' is at least warn.
func (logger Logger) Warn(format string, args ...interface{}) {
	if logger.Level >= WarnLevel {
		format = fmt.Sprintf("%5s  %s", "WARN", format)
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
		if logger.Date != "" {
			format = fmt.Sprintf("%s %s", time.Now().UTC().Format(logger.Date), format)
		}
		fmt.Fprintf(logger.Out, format, args...)
	}
}

// Error writes a log to the loggers 'Out' if its 'Level' is at least error.
func (logger Logger) Error(format string, args ...interface{}) {
	if logger.Level >= ErrorLevel {
		format = fmt.Sprintf("%5s  %s", "ERROR", format)
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
		if logger.Date != "" {
			format = fmt.Sprintf("%s %s", time.Now().UTC().Format(logger.Date), format)
		}
		fmt.Fprintf(logger.Out, format, args...)
	}
}

// Fatal writes a log to the loggers 'Out' if its 'Level' is at least fatal.
func (logger Logger) Fatal(format string, args ...interface{}) {
	if logger.Level >= FatalLevel {
		format = fmt.Sprintf("%5s  %s", "FATAL", format)
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
		if logger.Date != "" {
			format = fmt.Sprintf("%s %s", time.Now().UTC().Format(logger.Date), format)
		}
		fmt.Fprintf(logger.Out, format, args...)
	}
}

/*
Copyright 2017 Nanobox, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package log

type baseLogger struct {
	// The logging level the logger should log at. This is typically (and defaults
	// to) `Info`, which allows Info(), Warn(), Error() and Fatal() to be logged
	Level level

	// Embed a formatter
	Formatter

	// Embed a printer
	Printer
}

func newBaseLogger() *baseLogger {
	return &baseLogger{
		Level:     Info,
		Formatter: newBasicFormatter(),
		Printer:   newFmtPrinter(),
	}
}

// Set, change, or update the log level
func (logger *baseLogger) SetLevel(l level) {
	logger.Level = l
}

// print a TRACE message if the level is high enough
func (logger baseLogger) Trace(format string, args ...interface{}) {
	// ensure the level is at least trace
	if logger.Level < Trace {
		return
	}

	// format, and print it out
	logger.Print(logger.Format("TRACE", format), args...)
}

// print a DEBUG message if the level is high enough
func (logger baseLogger) Debug(format string, args ...interface{}) {
	// ensure the level is at least debug
	if logger.Level < Debug {
		return
	}

	// format, and print it out
	logger.Print(logger.Format("DEBUG", format), args...)
}

// print a INFO message if the level is high enough
func (logger baseLogger) Info(format string, args ...interface{}) {
	// ensure the level is at least info
	if logger.Level < Info {
		return
	}

	// format, and print it out
	logger.Print(logger.Format("INFO", format), args...)
}

// print a WARN message if the level is high enough
func (logger baseLogger) Warn(format string, args ...interface{}) {
	// ensure the level is at least warn
	if logger.Level < Warn {
		return
	}

	// format, and print it out
	logger.Print(logger.Format("WARN", format), args...)
}

// print a ERROR message if the level is high enough
func (logger baseLogger) Error(format string, args ...interface{}) {
	// ensure the level is at least error
	if logger.Level < Error {
		return
	}

	// format, and print it out
	logger.Print(logger.Format("ERROR", format), args...)
}

// print a FATAL message always
func (logger baseLogger) Fatal(format string, args ...interface{}) {
	// format, and print it out
	logger.Print(logger.Format("FATAL", format), args...)
}

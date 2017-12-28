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

type level uint32

type Logger interface {
	Opener
	Closer
	LevelSetter
	Tracer
	Debugger
	Infoer
	Warner
	Errorer
	Fataler
}

type Opener interface {
	Open() error
}

type Closer interface {
	Close() error
}

type LevelSetter interface {
	SetLevel(l level)
}

type Tracer interface {
	Trace(format string, args ...interface{})
}

type Debugger interface {
	Debug(format string, args ...interface{})
}

type Infoer interface {
	Info(format string, args ...interface{})
}

type Warner interface {
	Warn(format string, args ...interface{})
}

type Errorer interface {
	Error(format string, args ...interface{})
}

type Fataler interface {
	Fatal(format string, args ...interface{})
}

const (
	// Fatal level. Logs. Used for critical errors.
	Fatal level = iota
	// Error level. Logs. Used for errors that should definitely be noted.
	Error
	// Warn level. Non-critical entries that deserve eyes.
	Warn
	// Info level. General operational entries about what's going on inside the
	// application.
	Info
	// Debug level. Usually only enabled when debugging. Very verbose logging.
	Debug
	// Trace level. Usually only enabled when debugging. Very very verbose logging.
	Trace
)

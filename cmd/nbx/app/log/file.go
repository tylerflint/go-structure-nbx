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

import (
	"fmt"
	"os"
)

// let's create a interface scoped to exactly what we need
type FileWriterCloser interface {
	Write(b []byte) (n int, err error)
	Close() error
}

type FileLogger struct {
	Path string

	// open file handle
	file FileWriterCloser

	// function to open a file as an io.Writer
	fsOpenFunc func(name string, flag int, perm os.FileMode) (*os.File, error)

	// Embed a baseLogger
	*baseLogger
}

func NewFileLogger(path string) *FileLogger {
	return &FileLogger{
		Path:       path,
		fsOpenFunc: os.OpenFile,
		baseLogger: newBaseLogger(),
	}
}

func (logger *FileLogger) Open() error {
	file, err := logger.fsOpenFunc(logger.Path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return fmt.Errorf("Unable to open file (%s) for writing: %v", logger.Path, err)
	}

	// set the handle to be able to close later
	logger.file = file

	// create a new printer with the log handle as the out
	printer := newFmtPrinter()
	printer.Out = file
	logger.Printer = printer

	return nil
}

func (logger FileLogger) Close() error {
	if err := logger.file.Close(); err != nil {
		return fmt.Errorf("Failed to close file (%s): %v", logger.Path, err)
	}

	return nil
}

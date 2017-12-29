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
	"errors"
	"os"
	"reflect"
	"testing"
)

type FakeFileWriterCloser struct {
	err error
}

func (f FakeFileWriterCloser) Close() error {
	return f.err
}

func (f FakeFileWriterCloser) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func TestNewFileLogger(t *testing.T) {
	logger := NewFileLogger("/tmp")

	if logger == nil {
		t.Errorf("Logger should not be nil")
	}

	if logger.Path == "" {
		t.Errorf("Path should not be empty")
	}

	if logger.fsOpenFunc == nil {
		t.Errorf("Openfunc should not be nil")
	}

	if logger.baseLogger == nil {
		t.Errorf("baseLogger should not be nil")
	}
}

func TestFileOpen(t *testing.T) {
	errOpener := &FileLogger{
		Path: "/tmp",
		fsOpenFunc: func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return nil, errors.New("Bad permissions")
		},
	}

	if err := errOpener.Open(); err == nil {
		t.Errorf("Should bubble up an error")
	}

	file := &os.File{}
	okOpener := &FileLogger{
		Path: "/tmp",
		fsOpenFunc: func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return file, nil
		},
		baseLogger: newBaseLogger(),
	}

	err := okOpener.Open()
	if err != nil {
		t.Errorf("Was not expecting an error, got: %v", err)
	}

	if !reflect.DeepEqual(okOpener.file, file) {
		t.Errorf("File was not set properly")
	}
}

func TestFileClose(t *testing.T) {
	errCloser := &FileLogger{
		file: &FakeFileWriterCloser{
			err: errors.New("Bad filehandle"),
		},
	}

	okCloser := &FileLogger{
		file: &FakeFileWriterCloser{
			err: nil,
		},
	}

	if err := errCloser.Close(); err == nil {
		t.Errorf("Error should not be nil")
	}

	if err := okCloser.Close(); err != nil {
		t.Errorf("Was expecting err to be nil, got: %v", err)
	}
}

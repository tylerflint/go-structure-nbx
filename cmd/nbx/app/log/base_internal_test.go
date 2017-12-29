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
	"testing"
)

type FakePrinter struct {
	PrintRefs int
}

type FakeFormatter struct {
	FormatRefs int
}

func (p *FakePrinter) Print(format string, args ...interface{}) {
	p.PrintRefs += 1
}

func (f *FakeFormatter) Format(level, format string) string {
	f.FormatRefs += 1
	return fmt.Sprintf("%s %s", level, format)
}

func TestNewBaseLogger(t *testing.T) {
	logger := newBaseLogger()

	if logger == nil {
		t.Errorf("Logger should not be nil")
	}

	if logger.Level != Info {
		t.Errorf("Expecting default log level to be INFO")
	}

	if logger.Formatter == nil {
		t.Errorf("Formatter should not be nil")
	}

	if logger.Printer == nil {
		t.Errorf("Printer should not be nil")
	}
}

func TestSetLevel(t *testing.T) {
	logger := newBaseLogger()

	logger.SetLevel(Warn)

	if logger.Level != Warn {
		t.Errorf("Failed to set new log level")
	}
}

func TestBaseTrace(t *testing.T) {
	cases := []struct {
		logLevel           level
		expectedPrintRefs  int
		expectedFormatRefs int
	}{
		{
			logLevel:           Trace,
			expectedPrintRefs:  1,
			expectedFormatRefs: 1,
		},
		{
			logLevel:           Debug,
			expectedPrintRefs:  0,
			expectedFormatRefs: 0,
		},
	}

	for _, c := range cases {
		formatter := &FakeFormatter{}
		printer := &FakePrinter{}

		logger := &baseLogger{
			Level:     c.logLevel,
			Formatter: formatter,
			Printer:   printer,
		}

		logger.Trace("foo")

		if formatter.FormatRefs != c.expectedFormatRefs {
			t.Errorf("Expected %i calls to Format(), got: %i", c.expectedFormatRefs, formatter.FormatRefs)
		}

		if printer.PrintRefs != c.expectedPrintRefs {
			t.Errorf("Expected %i calls to Print(), got: %i", c.expectedPrintRefs, printer.PrintRefs)
		}
	}
}

func TestBaseDebug(t *testing.T) {
	cases := []struct {
		logLevel           level
		expectedPrintRefs  int
		expectedFormatRefs int
	}{
		{
			logLevel:           Debug,
			expectedPrintRefs:  1,
			expectedFormatRefs: 1,
		},
		{
			logLevel:           Info,
			expectedPrintRefs:  0,
			expectedFormatRefs: 0,
		},
	}

	for _, c := range cases {
		formatter := &FakeFormatter{}
		printer := &FakePrinter{}

		logger := &baseLogger{
			Level:     c.logLevel,
			Formatter: formatter,
			Printer:   printer,
		}

		logger.Debug("foo")

		if formatter.FormatRefs != c.expectedFormatRefs {
			t.Errorf("Expected %i calls to Format(), got: %i", c.expectedFormatRefs, formatter.FormatRefs)
		}

		if printer.PrintRefs != c.expectedPrintRefs {
			t.Errorf("Expected %i calls to Print(), got: %i", c.expectedPrintRefs, printer.PrintRefs)
		}
	}
}

func TestBaseInfo(t *testing.T) {
	cases := []struct {
		logLevel           level
		expectedPrintRefs  int
		expectedFormatRefs int
	}{
		{
			logLevel:           Info,
			expectedPrintRefs:  1,
			expectedFormatRefs: 1,
		},
		{
			logLevel:           Warn,
			expectedPrintRefs:  0,
			expectedFormatRefs: 0,
		},
	}

	for _, c := range cases {
		formatter := &FakeFormatter{}
		printer := &FakePrinter{}

		logger := &baseLogger{
			Level:     c.logLevel,
			Formatter: formatter,
			Printer:   printer,
		}

		logger.Info("foo")

		if formatter.FormatRefs != c.expectedFormatRefs {
			t.Errorf("Expected %i calls to Format(), got: %i", c.expectedFormatRefs, formatter.FormatRefs)
		}

		if printer.PrintRefs != c.expectedPrintRefs {
			t.Errorf("Expected %i calls to Print(), got: %i", c.expectedPrintRefs, printer.PrintRefs)
		}
	}
}

func TestBaseWarn(t *testing.T) {
	cases := []struct {
		logLevel           level
		expectedPrintRefs  int
		expectedFormatRefs int
	}{
		{
			logLevel:           Warn,
			expectedPrintRefs:  1,
			expectedFormatRefs: 1,
		},
		{
			logLevel:           Error,
			expectedPrintRefs:  0,
			expectedFormatRefs: 0,
		},
	}

	for _, c := range cases {
		formatter := &FakeFormatter{}
		printer := &FakePrinter{}

		logger := &baseLogger{
			Level:     c.logLevel,
			Formatter: formatter,
			Printer:   printer,
		}

		logger.Warn("foo")

		if formatter.FormatRefs != c.expectedFormatRefs {
			t.Errorf("Expected %i calls to Format(), got: %i", c.expectedFormatRefs, formatter.FormatRefs)
		}

		if printer.PrintRefs != c.expectedPrintRefs {
			t.Errorf("Expected %i calls to Print(), got: %i", c.expectedPrintRefs, printer.PrintRefs)
		}
	}
}

func TestBaseError(t *testing.T) {
	cases := []struct {
		logLevel           level
		expectedPrintRefs  int
		expectedFormatRefs int
	}{
		{
			logLevel:           Error,
			expectedPrintRefs:  1,
			expectedFormatRefs: 1,
		},
		{
			logLevel:           Fatal,
			expectedPrintRefs:  0,
			expectedFormatRefs: 0,
		},
	}

	for _, c := range cases {
		formatter := &FakeFormatter{}
		printer := &FakePrinter{}

		logger := &baseLogger{
			Level:     c.logLevel,
			Formatter: formatter,
			Printer:   printer,
		}

		logger.Error("foo")

		if formatter.FormatRefs != c.expectedFormatRefs {
			t.Errorf("Expected %i calls to Format(), got: %i", c.expectedFormatRefs, formatter.FormatRefs)
		}

		if printer.PrintRefs != c.expectedPrintRefs {
			t.Errorf("Expected %i calls to Print(), got: %i", c.expectedPrintRefs, printer.PrintRefs)
		}
	}
}

func TestBaseFatal(t *testing.T) {
	formatter := &FakeFormatter{}
	printer := &FakePrinter{}

	logger := &baseLogger{
		Level:     Fatal,
		Formatter: formatter,
		Printer:   printer,
	}

	logger.Fatal("foo")

	if formatter.FormatRefs != 1 {
		t.Errorf("Expected 1 call to Format(), got: %i", formatter.FormatRefs)
	}

	if printer.PrintRefs != 1 {
		t.Errorf("Expected 1 call to Print(), got: %i", printer.PrintRefs)
	}
}

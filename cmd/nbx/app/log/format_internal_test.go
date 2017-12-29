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
	"time"
)

func TestNewBasicFormatter(t *testing.T) {
	formatter := newBasicFormatter()

	if formatter == nil {
		t.Errorf("Formatter should not be nil")
	}

	if formatter.Date == "" {
		t.Errorf("Date should not be empty")
	}
}

func TestBasicFormat(t *testing.T) {
	cases := []struct {
		formatter basicFormatter
		level     string
		date      string
		msg       string
		newline   bool
	}{
		{
			formatter: basicFormatter{
				Date: "2006-01-02 15:04:05",
			},
			level:   "INFO",
			date:    time.Now().UTC().Format("2006-01-02 15:04:05"),
			msg:     "hello",
			newline: true,
		},
		{
			formatter: basicFormatter{
				Date: "2006-01-02 15:04:05",
			},
			level:   "INFO",
			date:    time.Now().UTC().Format("2006-01-02 15:04:05"),
			msg:     "hello",
			newline: false,
		},
		{
			formatter: basicFormatter{},
			level:     "INFO",
			date:      "",
			msg:       "hello",
			newline:   true,
		},
	}

	for _, c := range cases {
		var expected string

		if c.date != "" {
			expected = fmt.Sprintf("%s %5s  %s\n", c.date, c.level, c.msg)
		} else {
			expected = fmt.Sprintf("%5s  %s\n", c.level, c.msg)
		}

		var msg string

		if c.newline {
			msg = fmt.Sprintf("%s\n", c.msg)
		} else {
			msg = c.msg
		}

		actual := c.formatter.Format(c.level, msg)

		if actual != expected {
			t.Errorf("Expected '%s', got: '%s'", expected, actual)
		}
	}
}

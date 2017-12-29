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
	"time"
)

type Formatter interface {
	Format(level, format string) string
}

type basicFormatter struct {
	// The date format. If not set, timestamp is not printed.
	Date string
}

func newBasicFormatter() *basicFormatter {
	return &basicFormatter{
		Date: "2006-01-02 15:04:05.00000",
	}
}

func (b basicFormatter) Format(level, format string) string {
	// prefix the format with the level (TRACE, INFO, etc)
	format = fmt.Sprintf("%5s  %s", level, format)

	// Add a newline if it's missing
	if format[len(format)-1] != '\n' {
		format += "\n"
	}

	// add a date if configured
	if b.Date != "" {
		format = fmt.Sprintf("%s %s", time.Now().UTC().Format(b.Date), format)
	}

	return format
}

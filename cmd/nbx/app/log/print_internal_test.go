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
	"bytes"
	"testing"
)

func TestNewFmtPrinter(t *testing.T) {
	printer := newFmtPrinter()

	if printer == nil {
		t.Errorf("Printer should not be nil")
	}

	if printer.Out == nil {
		t.Errorf("Out should not be nil")
	}
}

func TestFmtPrint(t *testing.T) {
	var b bytes.Buffer

	printer := &fmtPrinter{
		Out: &b,
	}

	name := "Mickey"
	printer.Print("Hello %s", name)

	res := b.String()

	if res != "Hello Mickey" {
		t.Errorf("Expected 'Hello Mickey', got: %s", res)
	}
}

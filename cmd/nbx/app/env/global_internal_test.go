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

package env

import (
	"errors"
	"testing"
)

type FakeHomedirDetector struct {
	Dir string
	Err error
}

func (f FakeHomedirDetector) Homedir() (string, error) {
	return f.Dir, f.Err
}

func TestSystemGlobalDir(t *testing.T) {
	tmp := &FakeHomedirDetector{"/tmp/", nil}
	tmpGlobal := &SystemGlobalDirDetector{tmp}

	dir, err := tmpGlobal.GlobalDir()
	if err != nil {
		t.Errorf("Err should be nil, got: %v", err)
	}
	if dir != "/tmp/.nbx" {
		t.Errorf("Was expecting dir to be '/tmp/.nbx', got: %s", dir)
	}

	bad := &FakeHomedirDetector{"", errors.New("broken")}
	badGlobal := &SystemGlobalDirDetector{bad}

	dir, err = badGlobal.GlobalDir()
	if err == nil {
		t.Errorf("Was not expecting err to be nil")
	}
	if dir != "" {
		t.Errorf("Was expecting dir to be empty, got: %s", dir)
	}
}

func TestNewSystemGlobalDirDetector(t *testing.T) {
	detector := NewSystemGlobalDirDetector()

	if detector == nil {
		t.Errorf("SystemGlobalDirDetector{} should not be nil")
	}

	if detector.HomedirDetector == nil {
		t.Errorf("HomedirDetector should not be nil")
	}
}

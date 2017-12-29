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
	"os"
	"reflect"
	"testing"
)

type FakeGlobalDirDetector struct {
	Dir string
	Err error
}

func (f FakeGlobalDirDetector) GlobalDir() (string, error) {
	return f.Dir, f.Err
}

type FakeDirCreator struct {
	Err error
}

func (f FakeDirCreator) Mkdir(dir string, perms os.FileMode) error {
	return f.Err
}

func TestSystemBootstrap(t *testing.T) {
	cases := []struct {
		globalDirDetector GlobalDirDetector
		dirCreator        DirCreator
		expectedErr       error
	}{
		{
			globalDirDetector: FakeGlobalDirDetector{"/tmp", nil},
			dirCreator:        FakeDirCreator{nil},
			expectedErr:       nil,
		},
		{
			globalDirDetector: FakeGlobalDirDetector{
				Dir: "",
				Err: errors.New("something broke"),
			},
			dirCreator:  FakeDirCreator{nil},
			expectedErr: errors.New("Unable to determine global dir: something broke"),
		},
		{
			globalDirDetector: FakeGlobalDirDetector{"/tmp", nil},
			dirCreator:        FakeDirCreator{errors.New("bad permissions")},
			expectedErr:       errors.New("Failed to create dir (/tmp): bad permissions"),
		},
	}

	for _, c := range cases {
		bootstrapper := &SystemBootstraper{c.globalDirDetector, c.dirCreator}
		err := bootstrapper.Bootstrap()
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}
	}
}

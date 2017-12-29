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
	"fmt"
	"path/filepath"
)

type GlobalDirDetector interface {
	GlobalDir() (string, error)
}

type SystemGlobalDirDetector struct {
	HomedirDetector
}

func NewSystemGlobalDirDetector() *SystemGlobalDirDetector {
	return &SystemGlobalDirDetector{
		NewSystemHomedirDetector(),
	}
}

func (s SystemGlobalDirDetector) GlobalDir() (string, error) {
	// set Home based off the users homedir (~)
	path, err := s.Homedir()
	if err != nil {
		return "", fmt.Errorf("Unable to determine homedir: %v", err)
	}

	// generate the final path to ~/.nbx
	globalDir := filepath.ToSlash(filepath.Join(path, ".nbx"))

	return globalDir, nil
}

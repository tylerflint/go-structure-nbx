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
)

type Bootstraper interface {
  Bootstrap() error
}

type SystemBootstraper struct {
  GlobalDirDetector
  DirCreator
}

func NewSystemBootstraper() *SystemBootstraper {
  return &SystemBootstraper{
    NewSystemGlobalDirDetector(),
    NewSystemDirCreator(),
  }
}

// Bootstrap the app environment (global dir, permissions, log files, etc)
func (s SystemBootstraper) Bootstrap() error {
  globalDir, err := s.GlobalDir()
  if err != nil {
    return fmt.Errorf("Unable to determine global dir: %v", err)
  }
  
  if err := s.Mkdir(globalDir, 0755); err != nil {
    return fmt.Errorf("Failed to create dir (%s): %v", globalDir, err)
  }
  
  return nil
}

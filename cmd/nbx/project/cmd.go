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

package project

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/nbx/cmd/nbx/config"
)

func NewCommands() []*cobra.Command {
  return []*cobra.Command{
    newInitCommand(),
    newDestroyCommand(),
  }
}

func newInitCommand() *cobra.Command {
  return &cobra.Command{
    Use: "init",
    Short: "Initialize a local project",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("init")
    },
  }
}

func newDestroyCommand() *cobra.Command {
  return &cobra.Command{
    Use: "destroy",
    Short: "Destroy a local project",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("destroy")
    },
  }
}

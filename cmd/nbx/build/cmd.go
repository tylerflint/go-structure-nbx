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

package build

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	return []*cobra.Command{
		newBuildCommand(),
		newCompileCommand(),
		newPushCommand(),
	}
}

func newBuildCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Build the dev container image",
		Long:  "Build the dev container image",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("built!")
		},
	}
}

func newCompileCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "compile",
		Short: "Compile source and generate final container image",
		Long:  "Compile project source and generate the final container image",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("compiled!")
		},
	}
}

func newPushCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "push",
		Short: "Push the compiled image to a remote registry",
		Long:  "Push the compiled image to a remote registry",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pushed")
		},
	}
}

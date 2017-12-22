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

package registry

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	return []*cobra.Command{
		newRegistryCommand(),
	}
}

func newRegistryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registry",
		Short: "Remote docker registry management",
	}

	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newAddCommand())
	cmd.AddCommand(newRemoveCommand())

	return cmd
}

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List remote registries",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list")
		},
	}
}

func newAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a remote registry",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add")
		},
	}
}

func newRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "rm",
		Short: "Remove a remote registry",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remove")
		},
	}
}

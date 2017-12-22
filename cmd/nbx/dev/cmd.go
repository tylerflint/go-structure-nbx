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

package dev

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	return []*cobra.Command{
		newRunCommand(),
		newDevCommand(),
	}
}

func newRunCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run commands or a session within the dev sandbox",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deploy")
		},
	}
}

func newDevCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dev",
		Short: "Dev sandbox management",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("dev")
		},
	}

	cmd.AddCommand(newConsoleCommand())
	cmd.AddCommand(newTunnelCommand())

	return cmd
}

func newConsoleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "console",
		Short: "Open a console to a component within the dev sandbox",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("console")
		},
	}
}

func newTunnelCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tunnel",
		Short: "Open a tunnel to a component within the dev sandbox",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tunnel")
		},
	}
}

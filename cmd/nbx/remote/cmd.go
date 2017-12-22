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

package remote

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	return []*cobra.Command{
		newDeployCommand(),
		newRemoteCommand(),
	}
}

func newDeployCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploy changes to a remote app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deploy")
		},
	}
}

func newRemoteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remote",
		Short: "Manage remote apps",
	}

	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newAddCommand())
	cmd.AddCommand(newRemoveCommand())
	cmd.AddCommand(newConsoleCommand())
	cmd.AddCommand(newTunnelCommand())
	cmd.AddCommand(newLogCommand())

	return cmd
}

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List remote apps",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list")
		},
	}
}

func newAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a remote app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("app")
		},
	}
}

func newRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "rm",
		Short: "Remove a remote app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remove")
		},
	}
}

func newConsoleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "console",
		Short: "Open a console to a component within a remote app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("console")
		},
	}
}

func newTunnelCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tunnel",
		Short: "Open a tunnel to a component within a remote app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tunnel")
		},
	}
}

func newLogCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "logs",
		Short: "Stream logs from within a remote app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("logs")
		},
	}
}

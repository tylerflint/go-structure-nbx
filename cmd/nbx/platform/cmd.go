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

package platform

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	return []*cobra.Command{
		newStartCommand(),
		newStopCommand(),
		newStatusCommand(),
	}
}

func newStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the Nanobox platform",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("start")
		},
	}
}

func newStopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop the Nanobox platform",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("stop")
		},
	}
}

func newStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Check the status of the Nanobox platform",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("status")
		},
	}
}

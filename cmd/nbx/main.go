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

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	// "github.com/nanobox-io/nbx/cmd/nbx/app"
	"github.com/nanobox-io/nbx/cmd/nbx/boxfile"
	"github.com/nanobox-io/nbx/cmd/nbx/build"
	"github.com/nanobox-io/nbx/cmd/nbx/dev"
	"github.com/nanobox-io/nbx/cmd/nbx/display/message"
	"github.com/nanobox-io/nbx/cmd/nbx/dryrun"
	"github.com/nanobox-io/nbx/cmd/nbx/platform"
	"github.com/nanobox-io/nbx/cmd/nbx/project"
	"github.com/nanobox-io/nbx/cmd/nbx/registry"
	"github.com/nanobox-io/nbx/cmd/nbx/remote"
	"github.com/nanobox-io/nbx/cmd/nbx/user"
)

func main() {
	// parse config

	// setup logger

	// init db

	// create display

	// set the global app state

	// setup commands
	nbx := &cobra.Command{
		Use:   "nbx",
		Short: "nbx: Nanobox within your console",
		Long:  message.OverviewDescription(),
	}

	// add boxfile commands
	for _, cmd := range boxfile.NewCommands() {
		nbx.AddCommand(cmd)
	}

	// add build commands
	for _, cmd := range build.NewCommands() {
		nbx.AddCommand(cmd)
	}

	// add dev commands
	for _, cmd := range dev.NewCommands() {
		nbx.AddCommand(cmd)
	}

	// add dryrun commands
	for _, cmd := range dryrun.NewCommands() {
		nbx.AddCommand(cmd)
	}

	// add project commands
	for _, cmd := range project.NewCommands() {
		nbx.AddCommand(cmd)
	}
	//
	// add platform commands
	for _, cmd := range platform.NewCommands() {
		nbx.AddCommand(cmd)
	}

	// add registry commands
	for _, cmd := range registry.NewCommands() {
		nbx.AddCommand(cmd)
	}

	// add remote commands
	for _, cmd := range remote.NewCommands() {
		nbx.AddCommand(cmd)
	}

	// add user commands
	for _, cmd := range user.NewCommands() {
		nbx.AddCommand(cmd)
	}

	// run
	if err := nbx.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

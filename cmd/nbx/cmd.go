package main

import (
	"github.com/spf13/cobra"

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

func newNbxCommand() *cobra.Command {
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

	return nbx
}

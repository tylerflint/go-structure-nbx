package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nanobox-io/gila/cmd/nbx/boxfile"
	"github.com/nanobox-io/gila/cmd/nbx/config"
	"github.com/nanobox-io/gila/cmd/nbx/build"
	"github.com/nanobox-io/gila/cmd/nbx/dev"
	"github.com/nanobox-io/gila/cmd/nbx/display/message"
	"github.com/nanobox-io/gila/cmd/nbx/dryrun"
	"github.com/nanobox-io/gila/cmd/nbx/platform"
	"github.com/nanobox-io/gila/cmd/nbx/project"
	"github.com/nanobox-io/gila/cmd/nbx/registry"
	"github.com/nanobox-io/gila/cmd/nbx/remote"
	"github.com/nanobox-io/gila/cmd/nbx/user"
)

func main() {
  // parse config
	
  // setup logger
	
  // init db
	
  // create display

	// setup the global app state
	app := &config.App{
		Stdin: 	os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
    // Logger: logger,
    // DB: db,
    // Config: config
	}
	
  // setup commands
	nbx := &cobra.Command{
		Use:   	"nbx",
		Short: 	"nbx: Nanobox within your console",
		Long: 	message.OverviewDescription(),
	}
	
	// add boxfile commands
	for _, cmd := range boxfile.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	
	// add build commands
	for _, cmd := range build.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	
	// add dev commands
	for _, cmd := range dev.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	
	// add dryrun commands
	for _, cmd := range dryrun.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	
	// add project commands
	for _, cmd := range project.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	// 
	// add platform commands
	for _, cmd := range platform.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	
	// add registry commands
	for _, cmd := range registry.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	
	// add remote commands
	for _, cmd := range remote.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	
	// add user commands
	for _, cmd := range user.NewCommands(app) {
		nbx.AddCommand(cmd)
	}
	
  // run
	if err := nbx.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	
	os.Exit(0)
}

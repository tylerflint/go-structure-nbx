package dryrun

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newDryrunCommand(app),
  }
}

func newDryrunCommand(app *config.App) *cobra.Command {
  cmd := &cobra.Command{
    Use: "dry-run",
    Short: "Simulate a deploy",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("dry-run")
    },
  }
  
  cmd.AddCommand(newConsoleCommand(app))
  cmd.AddCommand(newTunnelCommand(app))
  
  return cmd
}

func newConsoleCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "console",
    Short: "Open a console to a component within the dry-run simulation",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("console")
    },
  }
}

func newTunnelCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "tunnel",
    Short: "Open a tunnel to a component within the dry-run simulation",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("tunnel")
    },
  }
}

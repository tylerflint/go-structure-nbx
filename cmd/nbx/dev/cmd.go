package dev

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newRunCommand(app),
    newDevCommand(app),
  }
}

func newRunCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "run",
    Short: "Run commands or a session within the dev sandbox",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("deploy")
    },
  }
}

func newDevCommand(app *config.App) *cobra.Command {
  cmd := &cobra.Command{
    Use: "dev",
    Short: "Dev sandbox management",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("dev")
    },
  }
  
  cmd.AddCommand(newConsoleCommand(app))
  cmd.AddCommand(newTunnelCommand(app))
  
  return cmd
}

func newConsoleCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "console",
    Short: "Open a console to a component within the dev sandbox",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("console")
    },
  }
}

func newTunnelCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "tunnel",
    Short: "Open a tunnel to a component within the dev sandbox",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("tunnel")
    },
  }
}

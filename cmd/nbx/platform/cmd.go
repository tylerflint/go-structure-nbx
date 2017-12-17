package platform

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newStartCommand(app),
    newStopCommand(app),
    newStatusCommand(app),
  }
}

func newStartCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "start",
    Short: "Start the Nanobox platform",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("start")
    },
  }
}

func newStopCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "stop",
    Short: "Stop the Nanobox platform",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("stop")
    },
  }
}

func newStatusCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "status",
    Short: "Check the status of the Nanobox platform",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("status")
    },
  }
}

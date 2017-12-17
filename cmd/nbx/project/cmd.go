package project

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newInitCommand(app),
    newDestroyCommand(app),
  }
}

func newInitCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "init",
    Short: "Initialize a local project",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("init")
    },
  }
}

func newDestroyCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "destroy",
    Short: "Destroy a local project",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("destroy")
    },
  }
}

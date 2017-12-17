package registry

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newRegistryCommand(app),
  }
}

func newRegistryCommand(app *config.App) *cobra.Command {
  cmd := &cobra.Command{
    Use: "registry",
    Short: "Remote docker registry management",
  }
  
  cmd.AddCommand(newListCommand(app))
  cmd.AddCommand(newAddCommand(app))
  cmd.AddCommand(newRemoveCommand(app))
  
  return cmd
}

func newListCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "ls",
    Short: "List remote registries",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("list")
    },
  }
}

func newAddCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "add",
    Short: "Add a remote registry",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("add")
    },
  }
}

func newRemoveCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "rm",
    Short: "Remove a remote registry",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("remove")
    },
  }
}

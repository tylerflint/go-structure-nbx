package remote

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newDeployCommand(app),
    newRemoteCommand(app),
  }
}

func newDeployCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "deploy",
    Short: "Deploy changes to a remote app",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("deploy")
    },
  }
}

func newRemoteCommand(app *config.App) *cobra.Command {
  cmd := &cobra.Command{
    Use: "remote",
    Short: "Manage remote apps",
  }
  
  cmd.AddCommand(newListCommand(app))
  cmd.AddCommand(newAddCommand(app))
  cmd.AddCommand(newRemoveCommand(app))
  cmd.AddCommand(newConsoleCommand(app))
  cmd.AddCommand(newTunnelCommand(app))
  cmd.AddCommand(newLogCommand(app))
  
  return cmd
}

func newListCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "list",
    Short: "List remote apps",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("list")
    },
  }
}

func newAddCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "add",
    Short: "Add a remote app",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("app")
    },
  }
}

func newRemoveCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "rm",
    Short: "Remove a remote app",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("remove")
    },
  }
}

func newConsoleCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "console",
    Short: "Open a console to a component within a remote app",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("console")
    },
  }
}

func newTunnelCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "tunnel",
    Short: "Open a tunnel to a component within a remote app",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("tunnel")
    },
  }
}

func newLogCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "logs",
    Short: "Stream logs from within a remote app",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("logs")
    },
  }
}

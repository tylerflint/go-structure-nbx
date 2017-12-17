package user

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newLoginCommand(app),
    newLogoutCommand(app),
  }
}

func newLoginCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "login",
    Short: "Login to the Nanobox platform",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("login")
    },
  }
}

func newLogoutCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "logout",
    Short: "Logout of the Nanobox platform",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("logout")
    },
  }
}

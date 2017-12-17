package boxfile

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newValidateCommand(app),
  }
}

func newValidateCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "validate",
    Short: "Validate the current boxfile.yml",
    Long: "Validate the current boxfile.yml",
    Run: func(cmd *cobra.Command, args []string) {
      
      if Validate("boxfile.yml") {
        fmt.Println("validated!")
      } else {
        fmt.Println("bad deal :(")
      }
      
    },
  }
}

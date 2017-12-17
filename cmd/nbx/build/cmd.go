package build

import (
  "fmt"
  
  "github.com/spf13/cobra"
  
  "github.com/nanobox-io/gila/cmd/nbx/config"
)

func NewCommands(app *config.App) []*cobra.Command {
  return []*cobra.Command{
    newBuildCommand(app),
    newCompileCommand(app),
    newPushCommand(app),
  }
}

func newBuildCommand(app *config.App) *cobra.Command {
  return &cobra.Command{
    Use: "build",
    Short: "Build the dev container image",
    Long: "Build the dev container image",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("built!")
    },
  }
}

func newCompileCommand(app *config.App) *cobra.Command {
  return &cobra.Command {
    Use: "compile",
    Short: "Compile source and generate final container image",
    Long: "Compile project source and generate the final container image",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("compiled!")
    },
  }
}

func newPushCommand(app *config.App) *cobra.Command {
  return &cobra.Command {
    Use: "push",
    Short: "Push the compiled image to a remote registry",
    Long: "Push the compiled image to a remote registry",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("pushed")
    },
  }
}

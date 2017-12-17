package config

import (
  "io"
)

type App struct {
  Stdin   io.Reader
  Stdout  io.Writer
  Stderr  io.Writer
  // Config
  // Display
  // Logger
  // DB
}

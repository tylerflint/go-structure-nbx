package app

import (
  "fmt"
  "os"
  "path/filepath"
  
  "github.com/mitchellh/go-homedir"
)

// Call dependencies in a mock-able way
var (
  homedirFunc   = homedir.Dir
  globalDirFunc = GlobalDir
  mkdirAllFunc  = os.MkdirAll
)

// Returns the global directory where Nanobox stores config, cache, logs, etc
func GlobalDir() (string, error) {

  // todo: allow an evar to set the final path

	// set Home based off the users homedir (~)
	path, err := homedirFunc()
	if err != nil {
		return "", fmt.Errorf("Unable to determine homedir: %v", err)
	}

  // generate the final path to ~/.nbx
	globalDir := filepath.ToSlash(filepath.Join(path, ".nbx"))

	return globalDir, nil
}

// Setup the app environment (global dir, permissions, log files, etc)
func SetupEnv() error {
  globalDir, err := globalDirFunc()
  if err != nil {
    return fmt.Errorf("Unable to determine global dir: %v", err)
  }
  
  if err := mkdirAllFunc(globalDir, 0755); err != nil {
    return fmt.Errorf("Failed to create dir (%s): %v", globalDir, err)
  }
  
  return nil
}

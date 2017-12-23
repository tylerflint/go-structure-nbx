package app

import (
  "errors"
  "os"
  "testing"
)

func TestGlobalDir(t *testing.T) {
  reset := homedirFunc
  
  homedirFunc = helperHomedirFuncTmp
  if globalDir, _ := GlobalDir(); globalDir != "/tmp/.nbx" {
    t.Errorf("Expected '/tmp/.nbx', got: %v", globalDir)
  }
  
  homedirFunc = helperHomedirFuncErr
  if _, err := GlobalDir(); err == nil {
    t.Errorf("Oops, should have been an error")
  }
  
  homedirFunc = reset
}

func TestSetupEnvOk(t *testing.T) {
  resetGlobalDir := globalDirFunc
  resetMkdirAll := mkdirAllFunc
  
  globalDirFunc = helperGlobalDirFuncTmp
  mkdirAllFunc = helperMkdirAllFuncOk
  
  if err := SetupEnv(); err != nil {
    t.Errorf("Wasn't expecting an error, got: %v", err)
  }
  
  globalDirFunc = resetGlobalDir
  mkdirAllFunc = resetMkdirAll
}

func TestSetupEnvFailGlobalDir(t *testing.T) {
  resetGlobalDir := globalDirFunc
  resetMkdirAll := mkdirAllFunc
  
  globalDirFunc = helperGlobalDirFuncErr
  mkdirAllFunc = helperMkdirAllFuncOk
  
  if err := SetupEnv(); err == nil {
    t.Errorf("Oops, I was expecting an error")
  }
  
  globalDirFunc = resetGlobalDir
  mkdirAllFunc = resetMkdirAll
}

func TestSetupEnvFailMkdirAll(t *testing.T) {
  resetGlobalDir := globalDirFunc
  resetMkdirAll := mkdirAllFunc
  
  globalDirFunc = helperGlobalDirFuncTmp
  mkdirAllFunc = helperMkdirAllFuncErr
  
  if err := SetupEnv(); err == nil {
    t.Errorf("Oops, I was expecting an error")
  }
  
  globalDirFunc = resetGlobalDir
  mkdirAllFunc = resetMkdirAll
}

// Helpers

func helperHomedirFuncTmp() (string, error) {
  return "/tmp/", nil
}

func helperHomedirFuncErr() (string, error) {
  return "", errors.New("Homedir failed")
}

func helperGlobalDirFuncTmp() (string, error) {
  return "/tmp/.nbx", nil
}

func helperGlobalDirFuncErr() (string, error) {
  return "", errors.New("Globaldir failed")
}

func helperMkdirAllFuncOk(dir string, perms os.FileMode) error {
  return nil
}

func helperMkdirAllFuncErr(dir string, perms os.FileMode) error {
  return errors.New("Failed to mkdir")
}

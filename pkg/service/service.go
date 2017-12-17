package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Service defines a service to be registered with the system.
type Service struct {
	name         string
	configFile   string
	config       string
	launchSystem string
	startCmd     []string
	stopCmd      []string
	statusCmd    []string
	removeCmd    []string
}

var (
	execCmd    = exec.Command
	execWhich  = exec.LookPath
	writeFile  = ioutil.WriteFile
	removeFile = os.RemoveAll
	initSystem string
)

func init() {
	// initialize the launch system var.
	initSystem = launchSystem()
}

// New creates and validates a new service object.
func New(name string, cmd ...string) (Service, error) {
	if len(cmd) == 0 {
		return Service{}, fmt.Errorf("No start command specified")
	}

	svc := Service{
		name:         name,
		configFile:   configFile(name),
		config:       config(name, cmd),
		launchSystem: initSystem,
		startCmd:     startCmd(name),
		stopCmd:      stopCmd(name),
		statusCmd:    statusCmd(name),
		removeCmd:    removeCmd(name),
	}

	return svc, svc.Validate()
}

// Validate validates whether a compatable launch system is found.
func (s Service) Validate() error {
	if s.launchSystem == "" {
		return fmt.Errorf("Compatable launch system not found!")
	}
	return nil
}

// Create registers a service with the system's service manager.
func (s Service) Create() error {
	// write service manifest
	return writeFile(s.configFile, []byte(s.config), 0600)
}

// Remove removes the service manifest from the system's service manager.
func (s Service) Remove() error {
	return removeFile(s.configFile)
}

// Stop stops the service using the system's service manager.
func (s Service) Stop() error {
	if len(s.stopCmd) < 2 {
		return fmt.Errorf("Bad stop command - '%s'", s.stopCmd)
	}
	return execCmd(s.stopCmd[0], s.stopCmd[1:]...).Run()
}

// Start starts the service using the system's service manager.
func (s Service) Start() error {
	if len(s.startCmd) < 2 {
		return fmt.Errorf("Bad start command - '%s'", s.startCmd)
	}
	return execCmd(s.startCmd[0], s.startCmd[1:]...).Run()
}

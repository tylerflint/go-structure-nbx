package service

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// TestNew tests the New() function.
func TestNew(t *testing.T) {
	New("nanobox-server", "nanobox", "server")
	New("nanobox-server")
	initSystem = ""
	New("nanobox-server", "nanobox", "server")
}

// TestCreate tests the Create() function.
func TestCreate(t *testing.T) {
	svc, _ := New("truer", "true")
	writeFile = helperWriteFilePass
	err := svc.Create()
	if err != nil {
		t.Errorf("Got error running Create() - %s", err.Error())
	}
	writeFile = helperWriteFileErr
	err = svc.Create()
	if err == nil {
		t.Errorf("Succeeded running Create()")
	}
}

// TestRemove tests the Remove() function.
func TestRemove(t *testing.T) {
	initSystem = "systemd"
	svc, _ := New("truer", "true")
	removeFile = helperRemoveAllPass
	err := svc.Remove()
	if err != nil {
		t.Errorf("Got error running Remove() - %s", err.Error())
	}
	removeFile = helperRemoveAllErr
	err = svc.Remove()
	if err == nil {
		t.Errorf("Succeeded running Remove()")
	}
}

// TestStop tests the Stop() function.
func TestStop(t *testing.T) {
	initSystem = "systemd"
	svc, _ := New("truer", "true")
	execCmd = helperCommandPass
	err := svc.Stop()
	if err != nil {
		t.Errorf("Got error running Stop() - %s", err.Error())
	}
	execCmd = helperCommandErr
	err = svc.Stop()
	if err == nil {
		t.Errorf("Succeeded running Stop()")
	}

	// test bad cmd length
	initSystem = ""
	svc, _ = New("truer", "true")
	execCmd = helperCommandErr
	err = svc.Stop()
	if err == nil {
		t.Errorf("Succeeded running bad len Stop() - %s", err.Error())
	}
}

// TestStart tests the Start() function.
func TestStart(t *testing.T) {
	initSystem = "systemd"
	svc, _ := New("truer", "true")
	execCmd = helperCommandPass
	err := svc.Start()
	if err != nil {
		t.Errorf("Got error running Start() - %s", err.Error())
	}
	execCmd = helperCommandErr
	err = svc.Start()
	if err == nil {
		t.Errorf("Succeeded running Start()")
	}

	// test bad cmd length
	initSystem = ""
	svc, _ = New("truer", "true")
	execCmd = helperCommandErr
	err = svc.Start()
	if err == nil {
		t.Errorf("Succeeded running bad len Start()")
	}
}

func helperWriteFilePass(filename string, data []byte, perm os.FileMode) error {
	return nil
}

func helperWriteFileErr(filename string, data []byte, perm os.FileMode) error {
	return fmt.Errorf("Fail")
}

func helperRemoveAllPass(filename string) error {
	return nil
}

func helperRemoveAllErr(filename string) error {
	return fmt.Errorf("Fail")
}

func helperCommandPass(cmd string, args ...string) *exec.Cmd {
	return exec.Command("true")
}

func helperCommandErr(cmd string, args ...string) *exec.Cmd {
	return exec.Command("false")
}

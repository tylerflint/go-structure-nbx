package log_test

import (
	"bytes"
	"testing"

	"github.com/nanobox-io/gila/pkg/log"
)

// TestLogging tests all the logging functions
func TestLogging(t *testing.T) {
	tmpOut := bytes.Buffer{}

	logger := log.New()
	logger.Out = &tmpOut
	logger.Level = log.TraceLevel
	logger.Date = ""

	logger.Printf("Prolly worked")
	if tmpOut.String() != "Prolly worked\n" {
		t.Errorf("Failed to Printf - '%s'", tmpOut.String())
	}
	tmpOut.Reset()

	logger.Trace("Prolly worked")
	if tmpOut.String() != "TRACE  Prolly worked\n" {
		t.Errorf("Failed to Printf - '%s'", tmpOut.String())
	}
	tmpOut.Reset()

	logger.Debug("Prolly worked")
	if tmpOut.String() != "DEBUG  Prolly worked\n" {
		t.Errorf("Failed to Printf - '%s'", tmpOut.String())
	}
	tmpOut.Reset()

	logger.Info("Prolly worked")
	if tmpOut.String() != " INFO  Prolly worked\n" {
		t.Errorf("Failed to Printf - '%s'", tmpOut.String())
	}
	tmpOut.Reset()

	logger.Warn("Prolly worked")
	if tmpOut.String() != " WARN  Prolly worked\n" {
		t.Errorf("Failed to Printf - '%s'", tmpOut.String())
	}
	tmpOut.Reset()

	logger.Error("Prolly worked")
	if tmpOut.String() != "ERROR  Prolly worked\n" {
		t.Errorf("Failed to Printf - '%s'", tmpOut.String())
	}
	tmpOut.Reset()

	logger.Fatal("Prolly worked")
	if tmpOut.String() != "FATAL  Prolly worked\n" {
		t.Errorf("Failed to Printf - '%s'", tmpOut.String())
	}
	tmpOut.Reset()

	logger.Date = "2006-01-02 15:04:05.00000"
	logger.Trace("Prolly worked")
	logger.Debug("Prolly worked")
	logger.Info("Prolly worked")
	logger.Warn("Prolly worked")
	logger.Error("Prolly worked")
	logger.Fatal("Prolly worked")
}

package service

import (
	"fmt"
	"reflect"
	"testing"
)

// simple string input/output
type sString struct {
	input    string
	expected string
}

// TestConfigFile tests the configFile() function
func TestConfigFile(t *testing.T) {
	systemdTests := []sString{
		{
			input:    "nanobox-server",
			expected: "/etc/systemd/system/nanobox-server.service",
		},
		{
			input:    "nanobox",
			expected: "/etc/systemd/system/nanobox.service",
		},
	}

	upstartTests := []sString{
		{
			input:    "nanobox-server",
			expected: "/etc/init/nanobox-server.conf",
		},
		{
			input:    "nanobox",
			expected: "/etc/init/nanobox.conf",
		},
	}

	initSystem = "systemd"
	for i := range systemdTests {
		out := configFile(systemdTests[i].input)
		if out != systemdTests[i].expected {
			t.Errorf("Failed to return correct systemd config path. Got '%s'", out)
		}
	}

	initSystem = "upstart"
	for i := range upstartTests {
		out := configFile(upstartTests[i].input)
		if out != upstartTests[i].expected {
			t.Errorf("Failed to return correct upstart config path. Got '%s'", out)
		}
	}

	initSystem = "fake"
	if out := configFile("nanobox-server"); out != "" {
		t.Errorf("Failed to return empty string. Got '%s'", out)
	}
}

// TestConfig tests the config() function.
func TestConfig(t *testing.T) {
	systemdExpected := `[Unit]
Description=nanobox-server
After=network.target

[Service]
Type=simple
ExecStart=nanobox server

[Install]
WantedBy=multi-user.target
`

	upstartExpected := `
script
nanobox server
end script`

	initSystem = "systemd"
	out := config("nanobox-server", []string{"nanobox", "server"})
	if out != systemdExpected {
		t.Errorf("Systemd config mismatch.")
		t.Logf("Got: %s", out)
	}

	initSystem = "upstart"
	out = config("nanobox-server", []string{"nanobox", "server"})
	if out != upstartExpected {
		t.Errorf("Upstart config mismatch.")
		t.Logf("Got: %s", out)
	}

	initSystem = "fake"
	out = config("nanobox-server", []string{"nanobox", "server"})
	if out != "" {
		t.Errorf("Fake config mismatch.")
		t.Logf("Got: %s", out)
	}
}

func helperWhichSysd(file string) (string, error) {
	if file != "systemctl" {
		return file, fmt.Errorf("Not found")
	}
	return file, nil
}

func helperWhichUps(file string) (string, error) {
	if file != "initctl" {
		return file, fmt.Errorf("Not found")
	}
	return file, nil
}

func helperWhichNone(file string) (string, error) {
	return "", fmt.Errorf("Not found")
}

// TestLaunchSystem tests the launchSystem() function.
func TestLaunchSystem(t *testing.T) {
	execWhich = helperWhichSysd
	out := launchSystem()
	if out != "systemd" {
		t.Errorf("Failed to match systemd. Got '%s'", out)
	}

	execWhich = helperWhichUps
	out = launchSystem()
	if out != "upstart" {
		t.Errorf("Failed to match upstart. Got '%s'", out)
	}

	execWhich = helperWhichNone
	out = launchSystem()
	if out != "" {
		t.Errorf("Failed to match none. Got '%s'", out)
	}
}

// TestStartCmd tests the startCmd() function.
func TestStartCmd(t *testing.T) {
	tests := []struct {
		init     string
		input    string
		expected []string
	}{
		{
			init:     "fake",
			input:    "nanobox-server",
			expected: nil,
		},
		{
			init:     "systemd",
			input:    "nanobox-server",
			expected: []string{"systemctl", "start", "nanobox-server.service"},
		},
		{
			init:     "upstart",
			input:    "nanobox-server",
			expected: []string{"initctl", "start", "nanobox-server"},
		},
	}

	for i := range tests {
		initSystem = tests[i].init
		out := startCmd(tests[i].input)
		if !reflect.DeepEqual(out, tests[i].expected) {
			t.Errorf("Failed to return correct %s start command. Got '%s'", tests[i].init, out)
		}
	}
}

// TestStopCmd tests the stopCmd() function.
func TestStopCmd(t *testing.T) {
	tests := []struct {
		init     string
		input    string
		expected []string
	}{
		{
			init:     "fake",
			input:    "nanobox-server",
			expected: nil,
		},
		{
			init:     "systemd",
			input:    "nanobox-server",
			expected: []string{"systemctl", "stop", "nanobox-server.service"},
		},
		{
			init:     "upstart",
			input:    "nanobox-server",
			expected: []string{"initctl", "stop", "nanobox-server"},
		},
	}

	for i := range tests {
		initSystem = tests[i].init
		out := stopCmd(tests[i].input)
		if !reflect.DeepEqual(out, tests[i].expected) {
			t.Errorf("Failed to return correct %s stop command. Got '%s'", tests[i].init, out)
		}
	}
}

// TestStatusCmd tests the statusCmd() function.
func TestStatusCmd(t *testing.T) {
	tests := []struct {
		init     string
		input    string
		expected []string
	}{
		{
			init:     "fake",
			input:    "nanobox-server",
			expected: nil,
		},
		{
			init:     "systemd",
			input:    "nanobox-server",
			expected: []string{"systemctl", "--no-pager", "status", "nanobox-server.service"},
		},
		{
			init:     "upstart",
			input:    "nanobox-server",
			expected: []string{"initctl", "status", "nanobox-server"},
		},
	}

	for i := range tests {
		initSystem = tests[i].init
		out := statusCmd(tests[i].input)
		if !reflect.DeepEqual(out, tests[i].expected) {
			t.Errorf("Failed to return correct %s status command. Got '%s'", tests[i].init, out)
		}
	}
}

// TestRemoveCmd tests the removeCmd() function.
func TestRemoveCmd(t *testing.T) {
	out := removeCmd("no-thing")
	if out != nil {
		t.Errorf("Failed to return correct remove command. Got '%s'", out)
	}
}

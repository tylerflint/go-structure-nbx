package service

import (
	"fmt"
	"strings"
)

// configFile returns the location to place the service "manifest".
func configFile(name string) string {
	fmtString := ""
	switch initSystem {
	case "systemd":
		fmtString = "/etc/systemd/system/%s.service"
	case "upstart":
		fmtString = "/etc/init/%s.conf"
	default:
		return ""
	}
	return fmt.Sprintf(fmtString, name)
}

// config returns the service "manifest" contents.
func config(name string, command []string) string {
	switch initSystem {
	case "systemd":

		// todo: in order to make this more flexible, use functional options
		// along with perhaps a template
		return fmt.Sprintf(`[Unit]
Description=%s
After=network.target

[Service]
Type=simple
ExecStart=%s

[Install]
WantedBy=multi-user.target
`, name, strings.Join(command, " "))

	// todo: double check upstart
	case "upstart":
		return fmt.Sprintf(`
script
%s
end script`, strings.Join(command, " "))
	}

	return ""
}

// initSystem returns the init system currently used.
func launchSystem() string {
	_, err := execWhich("systemctl")
	if err == nil {
		return "systemd"
	}

	_, err = execWhich("initctl")
	if err == nil {
		return "upstart"
	}

	return ""
}

// startCmd returns the init specific command to start the service.
func startCmd(name string) []string {
	switch initSystem {
	case "systemd":
		return []string{"systemctl", "start", fmt.Sprintf("%s.service", name)}
	case "upstart":
		return []string{"initctl", "start", name}
	}

	return nil
}

// stopCmd returns the init specific command to stop the service.
func stopCmd(name string) []string {
	switch initSystem {
	case "systemd":
		return []string{"systemctl", "stop", fmt.Sprintf("%s.service", name)}
	case "upstart":
		return []string{"initctl", "stop", name}
	}

	return nil
}

// statusCmd returns the init specific command to retrieve the status of the service.
func statusCmd(name string) []string {
	switch initSystem {
	case "systemd":
		return []string{"systemctl", "--no-pager", "status", fmt.Sprintf("%s.service", name)}
	case "upstart":
		return []string{"initctl", "status", name}
	}

	return nil
}

// removeCmd returns the init specific command to remove the service.
func removeCmd(name string) []string {
	return nil
}

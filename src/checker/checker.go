package checker

import (
	"fmt"
	"strings"

	"wolremote/src/ssh"
)

// CheckIfRunning checks if a given program is running on the remote machine
func CheckIfRunning(client *ssh.Client, programName string) (bool, error) {
	// Use pgrep -f which is cleaner for this purpose
	checkCmd := fmt.Sprintf("pgrep -f '%s'", programName)

	output, err := client.RunCommand(checkCmd)
	if err != nil {
		// If pgrep didn't find anything, it's okay
		if strings.Contains(err.Error(), "Process exited with status 1") {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if program is running: %w", err)
	}

	// If output is not empty, program is running
	if strings.TrimSpace(output) != "" {
		return true, nil
	}

	return false, nil
}

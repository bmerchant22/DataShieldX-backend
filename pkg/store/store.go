package store

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func StartCodeServer() error {
	// Command to start code-server
	cmd := exec.Command("code-server", "--bind-addr", "0.0.0.0:8081")

	// Set the command's stdout and stderr to the current process's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting code-server: %w", err)
	}

	// Wait for the command to finish
	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("code-server exited with error: %v\n", err)
		}
	}()

	return nil
}

func StopCodeServer() error {
	// Command to retrieve the PIDs of the running code-server processes
	pidCmd := exec.Command("pgrep", "-f", "code-server --bind-addr 0.0.0.0:8081")

	// Run the command and get the output
	output, err := pidCmd.Output()
	if err != nil {
		return fmt.Errorf("error retrieving code-server PIDs: %w", err)
	}

	// Split the output into individual PIDs
	pids := strings.Fields(string(output))

	// Loop through each PID and stop the corresponding code-server process
	for _, pidStr := range pids {
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			return fmt.Errorf("error converting PID to integer: %w", err)
		}

		// Command to stop code-server using its PID
		killCmd := exec.Command("kill", strconv.Itoa(pid))

		// Set the command's stdout and stderr to the current process's stdout and stderr
		killCmd.Stdout = os.Stdout
		killCmd.Stderr = os.Stderr

		// Run the command
		if err := killCmd.Run(); err != nil {
			return fmt.Errorf("error stopping code-server: %w", err)
		}
	}

	return nil
}

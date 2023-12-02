package store

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
)

var (
	codeServerPID int32
	codeServerLog atomic.Value
)

//func StartCodeServer() error {
//	// Command to start code-server
//	cmd := exec.Command("code-server", "--bind-addr", "0.0.0.0:8081")
//
//	// Set the command's stdout and stderr to the current process's stdout and stderr
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//
//	// Start the command
//	if err := cmd.Start(); err != nil {
//		return fmt.Errorf("error starting code-server: %w", err)
//	}
//
//	// Wait for the command to finish
//	go func() {
//		if err := cmd.Wait(); err != nil {
//			fmt.Printf("code-server exited with error: %v\n", err)
//		}
//	}()
//
//	return nil
//}
//
//func StopCodeServer() error {
//	// Command to retrieve the PIDs of the running code-server processes
//	pidCmd := exec.Command("pgrep", "-f", "code-server --bind-addr 0.0.0.0:8081")
//
//	// Run the command and get the output
//	output, err := pidCmd.Output()
//	if err != nil {
//		return fmt.Errorf("error retrieving code-server PIDs: %w", err)
//	}
//
//	// Split the output into individual PIDs
//	pids := strings.Fields(string(output))
//
//	// Loop through each PID and stop the corresponding code-server process
//	for _, pidStr := range pids {
//		pid, err := strconv.Atoi(pidStr)
//		if err != nil {
//			return fmt.Errorf("error converting PID to integer: %w", err)
//		}
//
//		// Command to stop code-server using its PID
//		killCmd := exec.Command("kill", strconv.Itoa(pid))
//
//		// Set the command's stdout and stderr to the current process's stdout and stderr
//		killCmd.Stdout = os.Stdout
//		killCmd.Stderr = os.Stderr
//
//		// Run the command
//		if err := killCmd.Run(); err != nil {
//			return fmt.Errorf("error stopping code-server: %w", err)
//		}
//	}
//
//	return nil
//}

func StartCodeServer() (string, error) {
	// Check if the server is already running
	if atomic.LoadInt32(&codeServerPID) != 0 {
		return "", fmt.Errorf("server is already running with PID %d", atomic.LoadInt32(&codeServerPID))
	}

	// Create a buffer to capture the logs
	logBuffer := &bytes.Buffer{}

	// Command to start code-server
	cmd := exec.Command("code-server", "--bind-addr", "0.0.0.0:8081")

	// Set the command's stdout and stderr to the buffer
	cmd.Stdout = logBuffer
	cmd.Stderr = logBuffer

	// Start the command
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("error starting code-server: %w", err)
	}

	// Get the PID of the code-server process
	atomic.StoreInt32(&codeServerPID, int32(cmd.Process.Pid))
	codeServerLog.Store(logBuffer)

	// Wait for the command to finish
	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("code-server exited with error: %v\n", err)
		}

		// Reset the saved PID and logs
		atomic.StoreInt32(&codeServerPID, 0)
		codeServerLog.Store((*bytes.Buffer)(nil))
	}()

	// Return the captured logs
	return logBuffer.String(), nil
}

func StopCodeServer() error {
	// Check if the server is not running
	if atomic.LoadInt32(&codeServerPID) == 0 {
		return fmt.Errorf("server is not running")
	}

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

	// Reset the saved PID and logs
	atomic.StoreInt32(&codeServerPID, 0)
	codeServerLog.Store((*bytes.Buffer)(nil))

	return nil
}

func GetLatestLogs() string {
	logBuffer, ok := codeServerLog.Load().(*bytes.Buffer)
	if !ok {
		return "No logs available"
	}

	return logBuffer.String()
}

func ResolvePath(dir string) (string, error) {
	home, err := os.UserHomeDir()

	if dir == "~" {
		if err != nil {
			return "", err
		}
		return home, nil
	}

	absPath := filepath.Join(home, dir)
	return absPath, nil
}

func GetDirDetails(dir string) ([]string, error) {
	// Read the files and folders in the directory
	var fileList []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fileList = append(fileList, file.Name())
	}

	return fileList, nil
}

func GetFileContent(absPath string) ([]byte, error) {
	content, err := os.ReadFile(absPath)
	if err != nil {
		zap.S().Errorf("Error while reading the file %v", absPath)
		return nil, nil
	}

	return content, err
}

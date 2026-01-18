// Package helpers provides utility functions for devops-scripts.
package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/google/uuid"
)

// GetRandomUUID generates a random UUID.
func GetRandomUUID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// GetOSName returns the name of the operating system.
func GetOSName() string {
	return runtime.GOOS
}

// GetOSVersion returns the version of the operating system.
func GetOSVersion() string {
	return runtime.Version()
}

// GetProcessName returns the name of the current process.
func GetProcessName() string {
	return os.Args[0]
}

// GetProcessID returns the ID of the current process.
func GetProcessID() int {
	return os.Getpid()
}

// GetCurrentDirectory returns the path of the current working directory.
func GetCurrentDirectory() string {
	return os.Getwd()
}

// GetHomeDirectory returns the path of the user's home directory.
func GetHomeDirectory() string {
	return os.Getenv("HOME")
}

// ExecuteShellCommand executes a shell command and returns the output.
func ExecuteShellCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command '%s': %v", command, err)
	}
	return string(output), nil
}

// ExecuteBinaryCommand executes a binary command and returns the output.
func ExecuteBinaryCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command '%s': %v", command, err)
	}
	return string(output), nil
}

// CopyFile copies the contents of one file to another.
func CopyFile(source, destination string) error {
	in, err := os.Open(source)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// DeleteFile deletes a file.
func DeleteFile(filename string) error {
	return os.Remove(filename)
}

// CreateDirectory creates a new directory.
func CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

// RemoveDirectory removes a directory.
func RemoveDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

// Exists checks if a file exists.
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// IsFile checks if a file is a file.
func IsFile(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

// IsDirectory checks if a file is a directory.
func IsDirectory(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// GetFileSize returns the size of a file.
func GetFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// GetFileModifiedTime returns the last modified time of a file.
func GetFileModifiedTime(filename string) (time.Time, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return time.Time{}, err
	}
	return fileInfo.ModTime(), nil
}

// GetFilePermissions returns the permissions of a file.
func GetFilePermissions(filename string) (os.FileMode, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Mode(), nil
}

// SetFilePermissions sets the permissions of a file.
func SetFilePermissions(filename string, permissions os.FileMode) error {
	return os.Chmod(filename, permissions)
}

// GetProcessInfo returns information about the current process.
func GetProcessInfo() (os.Proc, error) {
	return os.FindProcess(os.Getpid())
}

// KillProcess kills a process.
func KillProcess(pid int) error {
	return syscall.Kill(pid, syscall.SIGKILL)
}
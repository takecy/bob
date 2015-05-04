package cli

import (
	"fmt"
	"os"
	"os/exec"
)

// ExecCommand command
func ExecCommand(command string, args ...string) (err error) {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		Fatalf("error exec command: %s", err)
		return
	}

	return
}

// Fatalf logger
func Fatalf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "\n  %s\n\n", fmt.Sprintf(msg, args...))
	os.Exit(1)
}

// Printf Printf and exit
func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args)
	os.Exit(1)
}

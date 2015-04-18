package cli

import (
	"fmt"
	"os"
	"os/exec"
)

// Execute command
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

// Error logger
func Fatalf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "\n  %s\n\n", fmt.Sprintf(msg, args...))
	os.Exit(1)
}

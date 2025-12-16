package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Running ls -la command")

	// Just defining the command and not running
	successCmd := exec.Command("ls", "-la")

	// Running the command and saving output and err simultaneously
	output, err := successCmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Failed command execution: %s\n", err)
	} else {
		fmt.Printf("Output:\n%s\n", string(output))
	}

	// Failed command with exit codes
	failedCmd := exec.Command("ls", "non-existent-file")

	output2, err2 := failedCmd.CombinedOutput()

	if err2 != nil {
		fmt.Printf("Error: \n%s\n", err2)

		var exitErr *exec.ExitError

		if errors.As(err2, &exitErr) {
			exitCode := exitErr.ExitCode()

			fmt.Printf("Extracted exit code: %d\n", exitCode)
		}
	} else {
		fmt.Println("Command succeeded!", string(output2))
	}

}

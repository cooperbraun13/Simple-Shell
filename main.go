package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Keyboard is our standard input device (os.Stdin)
// Each time we press the enter key, a new line is created (\n)
// Everything written is stored in input
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		// Read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove newline character
	input = strings.TrimSuffix(input, "\n")

	// Prepare command to execute
	cmd := exec.Command(input)

	// Set correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error
	return cmd.Run()
}
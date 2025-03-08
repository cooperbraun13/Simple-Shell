package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
	"os/exec"
)

// Keyboard is our standard input device (os.Stdin)
// Each time we press the enter key, a new line is created (\n)
// Everything written is stored in input
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove newline character
	input = strings.TrimSuffix(input, '\n')

	// Prepare command to execute
	cmd := exec.Command(input)
}
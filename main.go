// Keyboard is our standard input device (os.Stdin)
// Each time we press the enter key, a new line is created (\n)
// Everything written is stored in input
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Read keyboard input
		input, err := reader.ReadString("\n")
		if err ! nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
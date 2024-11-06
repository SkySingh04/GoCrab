package filesys

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"go.uber.org/zap"
)

// Select an editor, prioritizing Vim, then Nano, then the system's default editor
func selectEditor() string {
	editor := "vim"
	if _, err := exec.LookPath(editor); err == nil {
		return editor
	}

	editor = "nano"
	if _, err := exec.LookPath(editor); err == nil {
		return editor
	}

	return os.Getenv("EDITOR")
}

// Open the selected editor to get Rust code as input, and return it as a string
func readRustCodeFromEditor(editor string, logger *zap.Logger) string {
	tmpFile := "temp_rust_code.rs"
	defer os.Remove(tmpFile)

	cmd := exec.Command(editor, tmpFile)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to open editor", zap.Error(err))
		os.Exit(1)
	}

	code, err := os.ReadFile(tmpFile)
	if err != nil {
		logger.Error("Failed to read input from editor", zap.Error(err))
		os.Exit(1)
	}

	return string(code)
}

// Read Rust code from standard input until ":q" is entered on a new line
func readRustCodeFromStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for scanner.Scan() {
		line := scanner.Text()
		if line == ":q" {
			break
		}
		input += line + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	return input
}
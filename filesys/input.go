package filesys

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// Main entry point for transpiling Rust code from an editor or stdin
func TranspileFromEditorOrInput(logger *zap.Logger, outputDir, outputFileName string) {
	editor := selectEditor()

	if editor == "" {
		fmt.Println("No editor available in PATH. Please enter Rust code below (type ':q' on a new line to finish):")
		rustCode := readRustCodeFromStdin()
		writeTranspiledCode(rustCode, logger, "", outputDir, outputFileName)
		return
	}

	rustCode := readRustCodeFromEditor(editor, logger)
	writeTranspiledCode(rustCode, logger, "", outputDir, outputFileName)
}

// Main entry point for transpiling Rust code from a file
func TranspileFromFile(logger *zap.Logger, rustFile, outputDir, outputFileName string) {
	rustCode, err := os.ReadFile(rustFile)
	if err != nil {
		logger.Error("Failed to read Rust file", zap.String("file", rustFile), zap.Error(err))
		os.Exit(1)
	}

	writeTranspiledCode(string(rustCode), logger, rustFile, outputDir, outputFileName)
}

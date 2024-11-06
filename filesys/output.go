package filesys

import (
	"GoCrab/transpiler"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

// Transpile Rust code and write the Go output to the specified directory and file
func writeTranspiledCode(rustCode string, logger *zap.Logger, inputFileName, outputDir, outputFileName string) {
	goCode, err := transpiler.Transpile(rustCode)
	if err != nil {
		logger.Error("Failed to transpile Rust code to Go code", zap.Error(err))
		os.Exit(1)
	}

	outputPath := prepareOutputFilePath(inputFileName, outputDir, outputFileName, logger)
	if err := os.WriteFile(outputPath, []byte(goCode), 0644); err != nil {
		logger.Error("Failed to write Go code to output file", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("Successfully transpiled Rust code to Go code", zap.String("output", outputPath))
}

// Prepare the output file path by ensuring the directory exists
func prepareOutputFilePath(inputFileName, outputDir, outputFileName string, logger *zap.Logger) string {
	if outputFileName == "" {
		if inputFileName == "" {
			outputFileName = "out.go"
		} else {
			outputFileName = strings.Split(inputFileName, ".rs")[0] + ".go"
		}
	}

	if outputDir == "" {
		outputDir = "."
	}

	outputPath := filepath.Join(outputDir, outputFileName)

	// Ensure the output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		logger.Info("Output directory does not exist. Creating directory.", zap.String("directory", outputDir))
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			logger.Error("Failed to create output directory", zap.Error(err))
			os.Exit(1)
		}
	}

	return outputPath
}

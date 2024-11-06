package main

import (
	"GoCrab/log"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	logger, err := log.New()
	if err != nil {
		fmt.Println("Failed to start the logger for the CLI", err)
		return
	}

	// Define the ASCII logo
	logo := `
		_____ _____ _____ ______  ___  ______ 
		|  __ \  _  /  __ \| ___ \/ _ \ | ___ \
		| |  \/ | | | /  \/| |_/ / /_\ \| |_/ /
		| | __| | | | |    |    /|  _  || ___ \
		| |_\ \ \_/ / \__/\| |\ \| | | || |_/ /
		\____/\___/ \____/\_| \_\_| |_/\____/ 
`
	fmt.Println(logo)
	var outputDir string
	var outputFileName string

	var rootCmd = &cobra.Command{
		Use:   "gocrab [file]",
		Short: "GoCrab is a Rust-to-Go transpiler",
		Long:  `GoCrab is a CLI tool to transpile Rust code to Go code.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rustFile := args[0]

			// Check if the file exists
			if _, err := os.Stat(rustFile); os.IsNotExist(err) {
				logger.Error("File not found", zap.String("file", rustFile))
				os.Exit(1)
			}

			// Set default output file name if not provided
			if outputFileName == "" {
				outputFileName = filepath.Base(rustFile)
				outputFileName = outputFileName[:len(outputFileName)-len(filepath.Ext(outputFileName))] + ".go"
			}

			// Set default output directory if not provided
			if outputDir == "" {
				outputDir = "."
			}

			outputPath := filepath.Join(outputDir, outputFileName)
			logger.Info("Transpiling from Rust to Go", zap.String("file", rustFile))
			logger.Info("Output file", zap.String("path", outputPath))

			// Call your transpilation logic here
		},
	}

	// Define flags for output directory and output file name
	rootCmd.Flags().StringVarP(&outputDir, "output-dir", "d", "", "Directory to save the output Go file")
	rootCmd.Flags().StringVarP(&outputFileName, "output-file", "o", "", "Name of the output Go file")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

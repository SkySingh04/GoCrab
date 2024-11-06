package main

import (
	"GoCrab/filesys"
	"GoCrab/log"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var outputDir, outputFileName string

func main() {
	logger, err := log.New()
	if err != nil {
		fmt.Println("Failed to start the logger", err)
		return
	}

	// ASCII logo
	logo := `
		_____ _____ _____ ______  ___  ______ 
		|  __ \  _  /  __ \| ___ \/ _ \ | ___ \
		| |  \/ | | | /  \/| |_/ / /_\ \| |_/ /
		| | __| | | | |    |    /|  _  || ___ \
		| |_\ \ \_/ / \__/\| |\ \| | | || |_/ /
		\____/\___/ \____/\_| \_\_| |_/\____/ 
`
	fmt.Println(logo)

	rootCmd := &cobra.Command{
		Use:   "gocrab [file]",
		Short: "GoCrab is a Rust-to-Go transpiler",
		Long:  `GoCrab is a CLI tool to transpile Rust code to Go code.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				filesys.TranspileFromEditorOrInput(logger, outputDir, outputFileName)
			} else {
				filesys.TranspileFromFile(logger, args[0], outputDir, outputFileName)
			}
		},
	}

	// Define flags
	rootCmd.Flags().StringVarP(&outputDir, "output-dir", "d", "", "Directory to save the output Go file")
	rootCmd.Flags().StringVarP(&outputFileName, "output-file", "o", "", "Name of the output Go file")

	// Execute CLI
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

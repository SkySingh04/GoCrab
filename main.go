package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cargotogo",
		Short: "CargoToGo is a Rust to Go transpiler",
		Long:  `CargoToGo is a CLI tool to transpile Rust code to Go code.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to CargoToGo!")
		},
	}

	var transpileCmd = &cobra.Command{
		Use:   "transpile [file]",
		Short: "Transpile a Rust file to Go",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			fmt.Printf("Transpiling %s from Rust to Go...\n", file)
		},
	}

	rootCmd.AddCommand(transpileCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskman",
	Short: "A simple CLI task manager",
	Long:  `taskman is a fast and simple command-line task manager built with Go.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(clearCmd)
}
// v0-0
// v2-1
// v5-1
// v8-1
// v10-1
// v13-0

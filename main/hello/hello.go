package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello [flags] MESSAGE",
	Short: "hello world sample",
	Long:  "display hello world message",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if *verbose {
			fmt.Println("PreRun")
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if *verbose {
			fmt.Println("PostRun")
		}
	},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello %s?\n", args[0])
	},
}

var verbose *bool

func init() {
	verbose = rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

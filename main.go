package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "squash"}
	var compressCmd = &cobra.Command{
		Use:   "compress",
		Short: "Compress a file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Compressing file...")
		},
	}

	rootCmd.AddCommand(compressCmd)
	rootCmd.Execute()
}

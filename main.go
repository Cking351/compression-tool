package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Init root command
	var rootCmd = &cobra.Command{
		Use:   "squash",
		Short: "Squash is a file compresion tool using Huffman and RLE encoding",
	}

	// Define compress cmd
	var compressCmd = &cobra.Command{
		Use:   "compress <filename> <encodingType>",
		Short: "compress a file with specified encoding (RLE, Huffman, RLE+huffman)",
		Run:   compressFile,
	}

	rootCmd.AddCommand(compressCmd)

	// Exec the CLI
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func compressFile(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: squash compress <filename> <encodingType>")
		return
	}

	// Get the filename and encoding type from the args
	fileName := args[0]
	encodingType := args[1]

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening fileL: %v\n", err)
		return
	}
	defer file.Close()

	// Read file content into byte slice
	data, err := readFile(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Compress
	compressedData, err := compress(data, encodingType)
	if err != nil {
		fmt.Printf("Error during compresion: %v\n", err)
		return
	}

	// Save compressed data to a new file
	outputFileName := fileName + ".compressed"
	err = os.WriteFile(outputFileName, compressedData, 0644)
	if err != nil {
		fmt.Printf("Error saving compressed file: %v\n", err)
		return
	}

	fmt.Printf("File succesfully compressing using the %s encoding method. Output saved to %s\n", encodingType, outputFileName)
}

package main

import (
	"fmt"
	"os"
)

func readFile(file *os.File) ([]byte, error) {
	// Read file into []byte
	data, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func compress(data []byte, encodingType string) ([]byte, error) {

	switch encodingType {
	case "RLE":
		err := runLengthEncode(&data)
		if err != nil {
			return nil, err
		}
	case "huffman":
		err := huffmanEncode(&data)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unsupported encoding type: %s", encodingType)
	}
	return data, nil
}

func runLengthEncode(data *[]byte) error {
	var encoded []byte

	count := 1
	for i := 1; i < len(*data); i++ {
		if (*data)[i] == (*data)[i-1] {
			count++
		} else {
			encoded = append(encoded, (*data)[i-1], byte(count))

			// Reset
			count = 1
		}
	}
	// Append the last char and its count
	encoded = append(encoded, (*data)[len(*data)-1], byte(count))
	*data = encoded
	return nil
}

func huffmanEncode(data *[]byte) error {

	// Build frequency map

	// Build huffman tree using freq map

	// Create huffman codes by traversing tree

	// Encode input

	return data, nil
}

package main

import "log"

func compress(data []byte, encodingType string) ([]byte, error) {
	if encodingType == "RLE" {
		err := runLength(&data)
		if err != nil {
			log.Fatal("There was an error: ", err)
		}

		if encodingType == "huffman" {
			err := huffman(&data)
			if err != nil {
				log.Fatal("There was an error: ", err)
			}
		}
	}
	return data, nil
}

func runLength(data *[]byte) error {
	return data, nil
}

func huffman(data *[]byte) error {
	return data, nil
}

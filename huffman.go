package main

import (
	"bytes"
	"sort"
)

type HuffmanNode struct {
	freq  int
	value byte
	left  *HuffmanNode
	right *HuffmanNode
	code  string
}

func createHuffmanTree(freq []int, values []byte) (*HuffmanNode, []string) {
	var nodes []*HuffmanNode
	codes := make([]string, len(values))

	for i, freqVal := range freq {
		nodes = append(nodes, &HuffmanNode{freq: freqVal, value: values[i]})
		codes[i] = ""
	}

	sort.Slice(nodes, func(i int, j int) bool { return nodes[i].freq < nodes[j].freq })

	for len(nodes) > 1 {
		min1 := nodes[0]
		min2 := nodes[1]

		if min1.freq > min2.freq {
			min1 = min2
			min2 = nodes[1]
		}

		parent := &HuffmanNode{left: min1, right: min2}
		parent.code = "0" + min1.code + "1" + min2.code
		parent.freq = min1.freq + min2.freq

		nodes = append(nodes, parent)
		sort.Slice(nodes, func(i int, j int) bool { return nodes[i].freq < nodes[j].freq })

		for i, node := range nodes {
			if node.left == nil && node.right == nil {
				codes[i] = parent.code + codes[i]
			}
		}
		nodes = nodes[:len(nodes)-1]
	}
	return nodes[0], codes
}

func encode(input []byte, huffmanTree *HuffmanNode) []byte {
	var buf bytes.Buffer
	for _, b := range input {
		bf := huffmanTree
		for bf != nil {
			buf.WriteByte(bf.code[0])
			if bf.value < b {
				bf = bf.right
			} else {
				bf = bf.left
			}
		}
	}
	return buf.Bytes()
}

func huffmanEncode(data *[]byte) error {

	// calc the frequency of of each byte
	freq := make([]int, 256)
	for _, b := range *data {
		freq[b]++
	}

	// Build huffman tree
	values := []byte{}
	for i := 0; i < 256; i++ {
		if freq[i] > 0 { // Only check bytes that appear at least once
			values = append(values, byte(i))
		}
	}

	// Create the huffman tree and generate codes
	_, codes := createHuffmanTree(freq, values)

	// Encode the input data
	var currentCode string
	for _, b := range *data {
		currentCode += codes[b]
	}

	// Create the buffer to hold the encoded bytes
	var buffer bytes.Buffer
	for i := 0; i < len(currentCode); i += 8 {
		end := i + 8
		if end < len(currentCode) {
			end = len(currentCode)
		}
		byteStr := currentCode[i:end]

		// Pad if less than 8 bits
		for len(byteStr) < 8 {
			byteStr = "0" + byteStr
		}

		// Convert the 8-bit string into a byte
		var byteValue byte
		for j := 0; j < 8; j++ {
			if byteStr[j] == '1' {
				byteValue |= 1 << (7 - j)
			}
		}
		buffer.WriteByte(byteValue)
	}
	*data = buffer.Bytes()
	return nil
}

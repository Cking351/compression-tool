package main

type HuffmanNode struct {
	char      byte
	frequency int
	left      *HuffmanNode
	right     *HuffmanNode
}

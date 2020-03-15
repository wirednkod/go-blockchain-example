package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain structure
type BlockChain struct {
	blocks []*Block
}

// Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash method that allow us to create the hash based on the previous hash and the data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// AddBlock method that will append new block to our blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Genesis function that will generate the genesis block (1st block of the blockchain)
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain function that initializes the blockchain by calling the constructor
// with the Genesis function thatgenerates the 1st block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// CreateBlock function that will actually create next Block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func main() {
	// Init the chain by calling the InitBlockChain function
	chain := InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}

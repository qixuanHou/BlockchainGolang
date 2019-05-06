package main

import (
	"fmt"
	"bytes"
	"crypto/sha256"
)
type BlockChain struct {
	blocks []*Block
}
type Block struct {
	Hash 		[] byte
	Data 		 []byte
	PrevHash []byte
}

func (b * Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash  = hash[:]
}

func CreateBlock(data string, PrevHash []byte) * Block {
	block := &Block{[]byte{}, []byte(data), PrevHash}
	block.DeriveHash()
	return block
}

func (chain * BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after gensis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("data in block: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}
}
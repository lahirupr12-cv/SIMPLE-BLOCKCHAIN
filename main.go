package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) addBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func initBlockchain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := initBlockchain()
	fmt.Println(chain)
	chain.addBlock("First block after genesis")
	fmt.Println(chain)
	chain.addBlock("second block after genesis")
	chain.addBlock("third block after genesis")

	// for _, block := range chain.blocks {
	// 	// fmt.Printf("previous hashed: %x\n", block.PrevHash)
	// 	// fmt.Printf("Data: %s\n", block.Data)
	// 	// fmt.Printf("hashed: %x\n", block.Hash)
	// }
}

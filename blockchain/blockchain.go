package blockchain

import (
	"github.com/raphaelrrcoelho/caboco-coin/block"
)

const dbFile = "blockchain.db"

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*block.Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewGenesisBlock creates the genesis Block
func NewGenesisBlock() *block.Block {
	return block.NewBlock("Bloco Primeiro", []byte{})
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{NewGenesisBlock()}}
}

package blockchain

import (
  "github.com/raphaelrrcoelho/caboco-coin/block"
)

type Blockchain struct {
  Blocks []*block.Block
}

func (bc *Blockchain) AddBlock(data string) {
  prevBlock := bc.Blocks[len(bc.Blocks) - 1]
  newBlock := block.NewBlock(data, prevBlock.Hash)

  bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock() *block.Block {
  return block.NewBlock("Bloco Primeiro", []byte{})
}

func NewBlockchain() *Blockchain {
  return &Blockchain{[]*block.Block{NewGenesisBlock()}}
}

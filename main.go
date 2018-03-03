package main

import (
  "fmt"
  "github.com/raphaelrrcoelho/caboco-coin/blockchain"
)

func main() {
  bc := blockchain.NewBlockchain()

  bc.AddBlock("Manda 1 CBC pro Caboco")
  bc.AddBlock("Manda mais 2 CBC pro Caboco")

  for _, block := range bc.Blocks {
    fmt.Printf("Hash Anterior: %x\n", block.PrevBlockHash)
    fmt.Printf("Dados: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)
    fmt.Println()
  }
}

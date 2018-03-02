package main

import (
  "fmt"
  "block"
  "blockchain"
)

func main() {
  bc := NewBlockchain()

  bc.AddBlock("Manda 1 CBC pro Caboco")
  bc.AddBlock("Manda mais 2 CBC pro Caboco")

  for _, block := range bc.blocks {
    fmt.Println("Previous Hash: %x\n", block.PrevBlockHash)
    fmt.Println("Data: %s\n", block.Data)
    fmt.Println("Hash: %x\n", block.Hash)
    fmt.Println()
  }
}

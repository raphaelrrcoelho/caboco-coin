package main

import (
	"fmt"
	"strconv"

	"github.com/raphaelrrcoelho/caboco-coin/blockchain"
	"github.com/raphaelrrcoelho/caboco-coin/proofofwork"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Manda 1 CBC pro Caboco")
	bc.AddBlock("Manda mais 2 CBC pro Caboco")

	fmt.Println("-----------------------")

	i := bc.NewIterator()

	for i.CurrentHash != nil {
		block := i.Next()
		fmt.Printf("Hash Anterior: %x\n", block.PrevBlockHash)
		fmt.Printf("Dados: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := proofofwork.NewProofOfWork(
			block.Timestamp,
			block.Data,
			block.PrevBlockHash,
			block.Nonce,
		)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

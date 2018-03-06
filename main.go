package main

import (
	"github.com/raphaelrrcoelho/caboco-coin/CLI"
	"github.com/raphaelrrcoelho/caboco-coin/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.DB.Close()

	cli := CLI.CLI{bc}
	cli.Run()
}

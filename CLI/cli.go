package CLI

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/raphaelrrcoelho/caboco-coin/blockchain"
	"github.com/raphaelrrcoelho/caboco-coin/proofofwork"
)

type CLI struct {
	BC *blockchain.Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlock(data)
	fmt.Println("Bloco adicionado com sucesso.")
}

func (cli *CLI) printChain() {
	bci := cli.BC.NewIterator()

	for bci.CurrentHash != nil {
		block := bci.Next()
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

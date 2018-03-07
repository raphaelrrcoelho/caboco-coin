package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	BC *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Utilização:")
	fmt.Println("  adicionabloco -dados DADOS - adiciona um bloco na blockchain")
	fmt.Println("  imprimechain - imprime todos os blocos da blockchain")
	fmt.Println()
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
	case "adicionabloco":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "imprimechain":
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

func (cli *CLI) addBlock(transactions []*Transaction) {
	cli.BC.AddBlock(transactions)
	fmt.Println("Bloco adicionado com sucesso.")
}

func (cli *CLI) printChain() {
	bci := cli.BC.NewIterator()

	for bci.CurrentHash != nil {
		block := bci.Next()
		fmt.Printf("Hash Anterior: %x\n", block.PrevBlockHash)
		fmt.Printf("Transações: %s\n", block.Transactions)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(
			block.Timestamp,
			block.Transactions,
			block.PrevBlockHash,
			block.Nonce,
		)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Utilização:")
	fmt.Println()
	fmt.Println(" conferesaldo -endereco ENDERECO (Confere saldo do ENDERECO)")
	fmt.Println(" criablockchain -endereco ENDERECO (Cria blockchain e envia a recompensa do bloco genesis para ENDERECO)")
	fmt.Println(" enviar -de DE -para PARA -montante MONTANTE (Envia MONTANTE de moedas do endereço DE para o endereço PARA)")
	fmt.Println(" imprimechain (Imprime todos os blocos da blockchain)")
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

	checkBalanceCmd := flag.NewFlagSet("checkBalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	checkBalanceData := checkBalanceCmd.String("endereco", "", "Endereço")
	createBlockchainData := createBlockchainCmd.String("endereco", "", "Endereço")
	sendFrom := sendCmd.String("de", "", "Endereço de origem")
	sendTo := sendCmd.String("para", "", "Endereço de destino")
	sendAmount := sendCmd.Int("montante", 0, "Montante a enviar")

	switch os.Args[1] {
	case "conferesaldo":
		err := checkBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "criablockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "enviar":
		err := sendCmd.Parse(os.Args[2:])
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

	if checkBalanceCmd.Parsed() {
		if *checkBalanceData == "" {
			checkBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.checkBalance(*checkBalanceData)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainData == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainData)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) createBlockchain(address string) {
	bc := CreateBlockchainDB(address)
	defer bc.DB.Close()
	fmt.Println("Blockchain criada com sucesso.")
}

func (cli *CLI) checkBalance(address string) {
	bc := NewBlockchain()
	defer bc.DB.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

func (cli *CLI) send(from string, to string, amount int) {
	bc := NewBlockchain()
	defer bc.DB.Close()
	fmt.Println("Blockchain criada com sucesso.")
}

func (cli *CLI) printChain() {
	//FIX ME: Remove need for create a empty blokchain
	bc := NewBlockchain()
	defer bc.DB.Close()

	bci := bc.NewIterator()

	for bci.CurrentHash != nil {
		block := bci.Next()
		fmt.Printf("Hash Anterior: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

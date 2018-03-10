package main

import (
	"fmt"
	"log"
	"os"

	bolt "github.com/coreos/bbolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "Papa-Chibé: O nascido no Pará"

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip []byte
	DB  *bolt.DB
}

// Iterator iterates thru the Blocks on the Blockchain
type Iterator struct {
	CurrentHash []byte
	db          *bolt.DB
}

// NewIterator creates a new Blockchain Iterator
func (bc *Blockchain) NewIterator() *Iterator {
	return &Iterator{bc.tip, bc.DB}
}

// Next returns next block starting from the tip
func (bci *Iterator) Next() *Block {
	var block *Block

	err := bci.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bucket.Get(bci.CurrentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bci.CurrentHash = block.PrevBlockHash

	return block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(transactions, lastHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err = bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

func dbExists() bool {
	_, err := os.Stat(dbFile)
	return !os.IsNotExist(err)
}

// NewBlockchain creates a new Blockchain
func NewBlockchain() *Blockchain {
	if dbExists() == false {
		fmt.Println("Nenhum banco de dados blockchain encontrado. Crie um antes.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		tip = bucket.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

// CreateBlockchainDB creates a new blockchain DB with genesis block
func CreateBlockchainDB(address string) *Blockchain {
	if dbExists() {
		fmt.Println("Bando de dados da Blockchain já existe")
		os.Exit(1)
	}

	fmt.Println("Nenhuma blockchain encontrada. Criando uma nova.")

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)

		bucket, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put([]byte("l"), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}

		tip = genesis.Hash

		return nil
	})

	return &Blockchain{tip, db}
}

package blockchain

import (
	"fmt"
	"log"

	bolt "github.com/coreos/bbolt"
	blk "github.com/raphaelrrcoelho/caboco-coin/block"
	tx "github.com/raphaelrrcoelho/caboco-coin/transaction"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinBaseData = "Papa-Chibé: O nascido no Pará"

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
func (i *Iterator) Next() *block.Block {
	var b *blk.Block

	err := i.db.View(func(dbtx *bolt.Tx) error {
		bucket := dbtx.Bucket([]byte(blocksBucket))
		encodedBlock := bucket.Get(i.CurrentHash)
		b = block.DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.CurrentHash = b.PrevBlockHash

	return b
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.DB.View(func(dbtx *bolt.Tx) error {
		bucket := dbtx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := blk.NewBlock(data, lastHash)

	err = bc.DB.Update(func(dbtx *bolt.Tx) error {
		bucket := dbtx.Bucket([]byte(blocksBucket))
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

// NewBlockchain creates a new Blockchain
func NewBlockchain(address string) *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(dbtx *bolt.Tx) error {
		bucket := dbtx.Bucket([]byte(blocksBucket))

		if bucket == nil {
			fmt.Println("Nenhuma blockchain encontrada. Criando uma nova.")
			cbtx := tx.NewCoinbaseTX(address, genesisCoinbaseData)
			genesis := blk.NewGenesisBlock(cbtx)

			bucket, err = dbtx.CreateBucket([]byte(blocksBucket))
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
		} else {
			tip = bucket.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

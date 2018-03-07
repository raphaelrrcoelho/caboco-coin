package block

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	pow "github.com/raphaelrrcoelho/caboco-coin/proofofwork"
	tx "github.com/raphaelrrcoelho/caboco-coin/transaction"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Transactions  []*tx.Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int64
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var b Block
	decoder := gob.NewDecoder(bytes.NewReader(d))

	err := decoder.Decode(&b)
	if err != nil {
		log.Panic(err)
	}

	return &b
}

// NewGenesisBlock creates the genesis Block
func NewGenesisBlock(coinbase *tx.Transaction) *Block {
	return NewBlock([]*tx.Transaction{coinbase}, []byte{})
}

// NewBlock creates and returns Block
func NewBlock(transactions []*tx.Transaction, prevBlockHash []byte) *Block {
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         int64(0),
	}

	proof := pow.NewProofOfWork(
		b.Timestamp,
		b.Transactions,
		b.PrevBlockHash,
		b.Nonce,
	)

	nonce, hash := proof.Run()
	b.Hash = hash[:]
	b.Nonce = nonce

	return b
}

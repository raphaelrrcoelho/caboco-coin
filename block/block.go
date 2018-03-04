package block

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/raphaelrrcoelho/caboco-coin/proofofwork"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Data          []byte
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
func NewGenesisBlock() *Block {
	return NewBlock("Bloco Primeiro", []byte{})
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         int64(0),
	}

	pow := proofofwork.NewProofOfWork(
		b.Timestamp,
		b.Data,
		b.PrevBlockHash,
		b.Nonce,
	)

	nonce, hash := pow.Run()
	b.Hash = hash[:]
	b.Nonce = nonce

	return b
}

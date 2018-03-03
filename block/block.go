package block

import (
	"time"
	"github.com/raphaelrrcoelho/caboco-coin/proofofwork"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce int64
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data: []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash: []byte{},
		Nonce: int64(0),
	}

	pow := proofofwork.NewProofOfWork(
		block.Timestamp,
		block.Data,
		block.PrevBlockHash,
		block.Nonce,
	)

	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

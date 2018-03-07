package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = int64(24)

var (
	maxNonce = int64(math.MaxInt64)
)

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	blockTimestamp    int64
	blockTransactions []*Transaction
	prevBlockHash     []byte
	target            *big.Int
	nonce             int64
}

func (pow *ProofOfWork) prepareData() []byte {
	data := bytes.Join(
		[][]byte{
			pow.prevBlockHash,
			[]byte{},
			IntToHex(pow.blockTimestamp),
			IntToHex(targetBits),
			IntToHex(pow.nonce),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte

	fmt.Printf(
		"Minerando o block contendo \"%s\"\n",
		[]byte{},
	)

	for pow.nonce < maxNonce {
		data := pow.prepareData()

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)

		hashInt.SetBytes(hash[:])

		// if hashInt is smaller than the target
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			pow.nonce++
		}
	}

	fmt.Printf("\n\n")
	return pow.nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData()
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	// if hashInt is smaller than the target
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(
	timestamp int64,
	transactions []*Transaction,
	prevBlockHash []byte,
	nonce int64,
) *ProofOfWork {
	// above  0fac49161af82ed938add1d8725835cc123a1a87b1b196488360e58d4bfb51e3
	// target 0000010000000000000000000000000000000000000000000000000000000000
	// below  0000008b0f41ec78bab747864db66bcb9fb89920ee75f43fdaaeb5544f7f76ca
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) // left shift by 256 - target

	pow := &ProofOfWork{timestamp, transactions, prevBlockHash, target, nonce}
	return pow
}
package proofofwork

import (
  "bytes"
  "crypto/sha256"
  "fmt"
  "math"
  "math/big"
  "github.com/raphaelrrcoelho/caboco-coin/utils"
)

const targetBits = 24

var (
  maxNonce = math.MaxInt64
)

type ProofOfWork struct {
  blockTimestamp int64
  blockData []byte
  prevBlockHash []byte
  target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
  data := bytes.Join(
    [][]byte{
      pow.prevBlockHash,
      pow.blockData,
      utils.IntToHex(int64(pow.blockTimestamp)),
      utils.IntToHex(int64(targetBits)),
      utils.IntToHex(int64(nonce)),
    },
    []byte{},
  )

  return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
  var hashInt big.Int
  var hash [32]byte
  nonce := 0

  fmt.Printf(
    "Minerando o block contendo \"%s\"\n",
    pow.blockData,
  )

  for nonce < maxNonce {
    data := pow.prepareData(nonce)

    hash = sha256.Sum256(data)
    fmt.Printf("\r%x", hash)

    hashInt.SetBytes(hash[:])

    // if hashInt is smaller than the target
    if hashInt.Cmp(pow.target) == -1 {
      break
    } else {
      nonce++
    }
  }

  fmt.Printf("\n\n")
  return nonce, hash[:]
}

func NewProofOfWork(
  timestamp int64,
  data []byte,
  prevBlockHash []byte,
  ) *ProofOfWork {
  // above  0fac49161af82ed938add1d8725835cc123a1a87b1b196488360e58d4bfb51e3
  // target 0000010000000000000000000000000000000000000000000000000000000000
  // below  0000008b0f41ec78bab747864db66bcb9fb89920ee75f43fdaaeb5544f7f76ca
  target := big.NewInt(1)
  target.Lsh(target, uint(256 - targetBits)) // left shift by 256 - target

  pow := &ProofOfWork{timestamp, data, prevBlockHash, target}
  return pow
}

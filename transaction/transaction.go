package transaction

import (
  "bytes"
  "encoding/gob"
  "fmt"
  "log"
  "crypto/sha256"
)

const subsidy = 10

// Transaction represents a CabocoCoin trasaction
type Transaction struct{
  ID []byte
  Vin []TXInput
  Vout []TXOutput
}

// TXInput represents a transaction input
type TXInput struct {
  Txid []byte
  Vout int
  ScriptSig string
}

// TXOutput represents a transaction output
type TXOutput struct {
  Value int
  ScriptPubKey string
}

// SetID sets ID of a transaction
func (tx *Transaction) SetID() {
  var encoded bytes.Buffer
  var hash [32]byte

  enc := gob.NewEncoder(&encoded)
  err := enc.Encode(tx)
  if err != nil {
    log.Panic(err)
  }

  hash = sha256.Sum256(encoded.Bytes())
  tx.ID = hash[:]
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
  if data == "" {
    data = fmt.Sprintf("Recompensa para '%s'", to)
  }

  txin := TXInput{[]byte{}, -1, data}
  txout := TXOutput{subsidy, to}

  tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
  tx.SetID()

  return &tx
}

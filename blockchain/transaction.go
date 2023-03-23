package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/realjf/blockchain_demo/helper"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	helper.Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{
		ID:      nil,
		Inputs:  []TxInput{txin},
		Outputs: []TxOutput{txout},
	}
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		helper.Handle(err)

		for _, out := range outs {
			input := TxInput{
				ID:  txID,
				Out: out,
				Sig: from,
			}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{
		Value:  amount,
		PubKey: to,
	})

	if acc > amount {
		outputs = append(outputs, TxOutput{
			Value:  acc - amount,
			PubKey: from,
		})
	}

	tx := Transaction{
		ID:      nil,
		Inputs:  inputs,
		Outputs: outputs,
	}
	tx.SetID()

	return &tx
}

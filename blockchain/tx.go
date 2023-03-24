package blockchain

import (
	"bytes"
	"encoding/gob"

	"github.com/realjf/blockchain_demo/helper"
	"github.com/realjf/blockchain_demo/wallet"
)

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

type TxOutputs struct {
	Outputs []TxOutput
}

func NewTxOutput(value int, address string) *TxOutput {
	txo := &TxOutput{Value: value, PubKeyHash: nil}
	txo.Lock([]byte(address))
	return txo
}

func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Equal(lockingHash, pubKeyHash)
}

func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := helper.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

func (outs TxOutputs) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(outs)
	helper.Handle(err)
	return buffer.Bytes()
}

func DeserializeOutputs(data []byte) TxOutputs {
	var outputs TxOutputs
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&outputs)
	helper.Handle(err)
	return outputs
}

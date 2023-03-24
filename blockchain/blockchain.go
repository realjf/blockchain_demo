package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"

	"github.com/realjf/blockchain_demo/helper"
)

const (
	dbPath      = "./tmp/blocks"
	dbFile      = "./tmp/blocks/MANIFEST"
	genesisData = "First Transaction from Genesis"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (chain *BlockChain) AddBlock(transactions []*Transaction) *Block {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		helper.Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = append(lastHash, val...)

			return nil
		})

		return err
	})

	helper.Handle(err)

	newBlock := NewBlock(transactions, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		helper.Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	helper.Handle(err)

	return newBlock
}

func Genesis(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func DBexists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}

func InitBlockChain(address string) *BlockChain {
	var lastHash []byte

	if DBexists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	helper.Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		cbtx := CoinbaseTx(address, genesisData)
		genesis := Genesis(cbtx)
		fmt.Println("Genesis created")
		err = txn.Set(genesis.Hash, genesis.Serialize())
		helper.Handle(err)
		err = txn.Set([]byte("lh"), genesis.Hash)

		lastHash = genesis.Hash

		return err
	})

	helper.Handle(err)
	blockchain := BlockChain{
		LastHash: lastHash,
		Database: db,
	}

	return &blockchain
}

func ContinueBlockChain(address string) *BlockChain {
	if !DBexists() {
		fmt.Println("No existing blockchain found, create one!")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	helper.Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		helper.Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = append(lastHash, val...)
			return nil
		})

		return err
	})
	helper.Handle(err)

	chain := BlockChain{
		LastHash: lastHash,
		Database: db,
	}

	return &chain
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{
		CurrentHash: chain.LastHash,
		Database:    chain.Database,
	}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		helper.Handle(err)
		var encodedBlock []byte
		err = item.Value(func(val []byte) error {
			encodedBlock = append(encodedBlock, val...)
			return nil
		})

		block = Deserialize(encodedBlock)

		return err
	})
	helper.Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
}

// func (chain *BlockChain) FindUnspentTransactions(pubKeyHash []byte) []Transaction {
// 	var unspentTxs []Transaction

// 	spentTXOs := make(map[string][]int)

// 	iter := chain.Iterator()

// 	for {
// 		block := iter.Next()

// 		for _, tx := range block.Transactions {
// 			txID := hex.EncodeToString(tx.ID)

// 		Outputs:
// 			for outIdx, out := range tx.Outputs {
// 				if spentTXOs[txID] != nil {
// 					for _, spentOut := range spentTXOs[txID] {
// 						if spentOut == outIdx {
// 							continue Outputs
// 						}
// 					}
// 				}

// 				if out.IsLockedWithKey(pubKeyHash) {
// 					unspentTxs = append(unspentTxs, *tx)
// 				}

// 				if !tx.IsCoinbase() {
// 					for _, in := range tx.Inputs {
// 						if in.UsesKey(pubKeyHash) {
// 							inTxID := hex.EncodeToString(in.ID)
// 							spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Out)
// 						}
// 					}
// 				}
// 			}
// 		}

// 		if len(block.PrevHash) == 0 {
// 			break
// 		}
// 	}

// 	return unspentTxs
// }

func (chain *BlockChain) FindUTXO() map[string]TxOutputs {
	UTXO := make(map[string]TxOutputs)
	spentTXOs := make(map[string][]int)
	iter := chain.Iterator()

	for {
		block := iter.Next()
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Outputs {
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
			}
			if !tx.IsCoinbase() {
				for _, in := range tx.Inputs {
					inTxID := hex.EncodeToString(in.ID)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Out)
				}
			}
		}
		if len(block.PrevHash) == 0 {
			break
		}
	}

	return UTXO
}

// func (chain *BlockChain) FindSpendableOutputs(pubKeyHash []byte, amount int) (int, map[string][]int) {
// 	unspentOuts := make(map[string][]int)
// 	unspentTxs := chain.FindUnspentTransactions(pubKeyHash)
// 	accumulated := 0

// Work:
// 	for _, tx := range unspentTxs {
// 		txID := hex.EncodeToString(tx.ID)

// 		for outIdx, out := range tx.Outputs {
// 			if out.IsLockedWithKey(pubKeyHash) && accumulated < amount {
// 				accumulated += out.Value
// 				unspentOuts[txID] = append(unspentOuts[txID], outIdx)

// 				if accumulated >= amount {
// 					break Work
// 				}
// 			}
// 		}
// 	}

// 	return accumulated, unspentOuts
// }

func (chain *BlockChain) FindTransaction(ID []byte) (Transaction, error) {
	iter := chain.Iterator()

	for {
		block := iter.Next()

		for _, tx := range block.Transactions {
			if bytes.Equal(tx.ID, ID) {
				return *tx, nil
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}

	return Transaction{}, errors.New("Transaction does not exist")
}

func (chain *BlockChain) SignTransaction(tx *Transaction, privateKey *ecdsa.PrivateKey) {
	prevTXs := make(map[string]Transaction)

	for _, in := range tx.Inputs {
		prevTX, err := chain.FindTransaction(in.ID)
		helper.Handle(err)
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}
	tx.Sign(privateKey, prevTXs)
}

func (chain *BlockChain) VerifyTransaction(tx *Transaction) bool {
	prevTXs := make(map[string]Transaction)

	for _, in := range tx.Inputs {
		prevTX, err := chain.FindTransaction(in.ID)
		helper.Handle(err)
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	return tx.Verify(prevTXs)
}

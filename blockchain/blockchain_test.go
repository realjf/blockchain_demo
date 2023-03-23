package blockchain_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/realjf/blockchain_demo/blockchain"
)

func TestBlockChain(t *testing.T) {
	cases := map[string]struct {
		f func()
	}{
		"init": {
			f: func() {
				chain := blockchain.InitBlockChain()
				chain.AddBlock("First Block after Genesis")
				chain.AddBlock("Second Block after Genesis")
				chain.AddBlock("Third Block after Genesis")

				for _, block := range chain.GetBlocks() {
					t.Logf("Previous Hash: %x\n", block.PrevHash)
					t.Logf("Data in Block: %s\n", block.Data)
					t.Logf("Hash: %x\n", block.Hash)
					t.Logf("Nonce: %d\n", block.Nonce)

					pow := blockchain.NewProof(block)
					fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
					fmt.Println()
				}
			},
		},
	}

	for name, ts := range cases {
		t.Run(name, func(t *testing.T) {
			ts.f()
		})
	}
}

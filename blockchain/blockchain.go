package blockchain

type BlockChain struct {
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := NewBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func (chain *BlockChain) Len() int { return len(chain.blocks) }

func (chain *BlockChain) GetBlocks() []*Block {
	return chain.blocks
}

func genesis() *Block {
	return NewBlock("Genesis", []byte{})
}

func newBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*Block{genesis()},
	}
}

func InitBlockChain() *BlockChain {
	return newBlockChain()
}

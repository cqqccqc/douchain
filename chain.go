package main

// BlockChain block chain
type Chain struct {
	Blocks []*Block
}

// AddBlock add a block
func (bc *Chain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Header.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewGenesisBlock create genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockChain create block chain
func NewBlockChain() *Chain {
	return &Chain{[]*Block{NewGenesisBlock()}}
}

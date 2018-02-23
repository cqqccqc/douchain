package main

import (
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain_%s.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// Chain block chain
type Chain struct {
	tip []byte
	db  *bolt.DB
}

// AddBlock add a block
func (bc *Chain) AddBlock(data string) {

}

// NewGenesisBlock create genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockChain create block chain
func NewBlockChain() *Chain {
	var tip []byte
	db, _ := bolt.Open(dbFile, 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesis := NewGenesisBlock()
			b, _ := tx.CreateBucket([]byte(blocksBucket))
			b.Put(genesis.Header.Hash, genesis.Serialize())
			b.Put([]byte("l"), genesis.Header.Hash)
			tip = genesis.Header.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	bc := Chain{tip, db}
	return &bc
}

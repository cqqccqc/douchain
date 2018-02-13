package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block represent a block in block chain
type Block struct {
	Header *Header
	Data   []byte
}

func (block *Block) generateHash()  {
	timestamp := []byte(strconv.FormatInt(block.Header.Timestamp, 10))
	meta := bytes.Join([][]byte{block.Header.PrevBlockHash, block.Data, timestamp}, []byte{})

	hash := sha256.Sum256(meta)
	block.Header.Hash = hash[:]
}

// NewBlock create block
func NewBlock(data string, prevBlockHash []byte) *Block {
	header := &Header{
		Timestamp:time.Now().Unix(),
		PrevBlockHash:prevBlockHash,
	}
	block := &Block{
		Header: header,
		Data: []byte(data),
	}
	block.generateHash()
	return block
}
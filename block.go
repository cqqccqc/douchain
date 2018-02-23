package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
)

// Block represent a block in block chain
type Block struct {
	Header *Header
	Data   []byte
}

func (block *Block) generateHash() {
	timestamp := []byte(strconv.FormatInt(block.Header.Timestamp, 10))
	meta := bytes.Join([][]byte{block.Header.PrevBlockHash, block.Data, timestamp}, []byte{})

	hash := sha256.Sum256(meta)
	block.Header.Hash = hash[:]
}

// Serialize serialize block
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(block)
	return result.Bytes()
}

// DeserializeBlock deserialize block
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&block)
	return &block
}

// NewBlock create block
func NewBlock(data string, prevBlockHash []byte) *Block {
	header := &Header{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
	}
	block := &Block{
		Header: header,
		Data:   []byte(data),
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Work()

	// block.generateHash()
	block.Header.Hash = hash
	block.Header.Nonce = nonce
	return block
}

package main

import (
	"fmt"
	"crypto/sha256"
	"bytes"
)

func main() {


	bc := NewBlockChain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, b := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", b.Header.PrevBlockHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Header.Hash)
		fmt.Println()
	}

	strByte := []byte("I like donuts")
	allByte := bytes.Join([][]byte{
		strByte,
		[]byte(),
	}, []byte{})
	hash := sha256.Sum256([]byte("I like donuts"))
	fmt.Println(hash)
}

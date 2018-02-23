package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {

	strByte := []byte("I like donuts")
	allByte := bytes.Join([][]byte{
		strByte,
		[]byte("123"),
	}, []byte{})

	hash := sha256.Sum256(allByte)
	fmt.Println(hash)
}

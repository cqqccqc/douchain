package main

// Header block header
type Header struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

package main

import "math/big"

const targetBits = 24

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

func NewProofOfWork(b *Block)  *ProofOfWork{

	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))

	pow := &ProofOfWork{
		b,
		target,
	}

	return pow
}

func (pow *ProofOfWork) Work()  {
	
}
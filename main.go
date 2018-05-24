package main

func main() {
	// bc := NewBlockchain()
	// bc.AddBlock("Send 1 BTC to Ivan")
	// bc.AddBlock("Send 2 more BTC to Ivan")

	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	pow := NewProofOfWork(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

	// data1 := []byte("I like donuts")
	// data2 := []byte("I like donutsca07ca")
	// targetBits := 24
	// target := big.NewInt(1)
	// target.Lsh(target, uint(256-targetBits))
	// fmt.Printf("%x\n", sha256.Sum256(data1))
	// fmt.Printf("%64x\n", target)
	// fmt.Printf("%x\n", sha256.Sum256(data2))

	bc := NewBlockchain()

	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}

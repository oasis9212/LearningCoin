package main

import (
	"fmt"
	"ralo/blockchain"
)

func main() {

	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.AllBlocks() {
		fmt.Printf("[Data : %s]\n", block.Data)
		fmt.Printf("[Hash : %s]\n", block.Hash)
		fmt.Printf("[PrevHash : %s]\n", block.PrevHash)
	}
}

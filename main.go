package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string //  저장할 내용
	hash     string // 저장할 해쉬 내용.
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
} // 마지막 블록에 대한 해쉬값가져오기.

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newBlock)
}
func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("data is %s\n", block.data)
		fmt.Printf("data hash is %s\n", block.hash)
		fmt.Printf("data prevhash is %s\n", block.prevHash)
	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}

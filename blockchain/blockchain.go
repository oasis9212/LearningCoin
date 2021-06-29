package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// 싱글통 패턴.
type block struct {
	Data     string //  저장할 내용
	Hash     string // 저장할 해쉬 내용.
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) claculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)

}

func getPrevHash() string {
	totalBlock := len(GetBlockChain().blocks)
	if totalBlock == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlock-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getPrevHash()}
	newBlock.claculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *blockchain {
	if b == nil {
		once.Do(func() { //딱한번만 실행시키는 방법.
			b = &blockchain{}
			b.AddBlock("Genesis")

		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return GetBlockChain().blocks
}

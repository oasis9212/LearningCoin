package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// 싱글통 패턴.
type block struct {
	data     string //  저장할 내용
	hash     string // 저장할 해쉬 내용.
	prevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) claculateHash() string {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)

}

func getPrevHash() string {
	totalBlock := len(GetBlockChain().blocks)
	if totalBlock == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlock-1].hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getPrevHash()}
	newBlock.claculateHash()
	return &newBlock
}

func GetBlockChain() *blockchain {
	if b == nil {
		once.Do(func() { //딱한번만 실행시키는 방법.
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b
}

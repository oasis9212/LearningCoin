package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

// 싱글통 패턴.
type Block struct {
	Data     string `json:"data"` //  저장할 내용
	Hash     string `json:"hash"` // 저장할 해쉬 내용.
	PrevHash string `json:"prev_hash,omitempty"`
	Height   int    `json:"height"`
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) claculateHash() {
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

func createBlock(data string) *Block {
	newBlock := Block{data, "", getPrevHash(), len(GetBlockChain().blocks) + 1}
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

func (b *blockchain) AllBlocks() []*Block {
	return GetBlockChain().blocks
}

var ErrNotFound = errors.New("block not found")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil
}

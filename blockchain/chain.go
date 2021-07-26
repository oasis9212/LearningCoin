package blockchain

import (
	"errors"
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

func (b *blockchain) Addblcok(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

var b *blockchain
var once sync.Once

func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() { //딱한번만 실행시키는 방법.
			b = &blockchain{"", 0}
			b.Addblcok("Genesis")

		})
	}
	return b
}

var ErrNotFound = errors.New("block not found")

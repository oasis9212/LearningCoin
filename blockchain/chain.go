package blockchain

import (
	"fmt"
	"ralo/db"
	"ralo/utils"
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

func (b *blockchain) Addblcok(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash

	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks

}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.Tobytes(b))
}

func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() { //딱한번만 실행시키는 방법.
			b = &blockchain{"", 0}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.Addblcok("Genesis")
			} else {

				b.restore(checkpoint)
			}

		})
	}

	fmt.Println(b.NewestHash)
	return b
}

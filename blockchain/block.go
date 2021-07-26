package blockchain

import (
	"crypto/sha256"
	"fmt"
)

// 싱글통 패턴.
type Block struct {
	Data     string `json:"data"` //  저장할 내용
	Hash     string `json:"hash"` // 저장할 해쉬 내용.
	PrevHash string `json:"prev_hash,omitempty"`
	Height   int    `json:"height"`
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		data,
		"",
		prevHash,
		height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%s", sha256.Sum256([]byte(payload)))
	return block
}

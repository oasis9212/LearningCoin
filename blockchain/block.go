package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"ralo/db"
	"ralo/utils"
	"strings"
)

const difficulty int = 2

// 싱글통 패턴.
type Block struct {
	Data       string `json:"data"` //  저장할 내용
	Hash       string `json:"hash"` // 저장할 해쉬 내용.
	PrevHash   string `json:"prev_hash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

var ErrNotFound = errors.New(("block not found"))

func (b *Block) Persist() {
	db.SaveBlock(b.Hash, utils.Tobytes(b))
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		blockAsString := fmt.Sprint(b)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprint(b))))
		fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", blockAsString, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}

}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		data,
		"",
		prevHash,
		height,
		difficulty,
		0,
	}
	block.mine()
	block.Persist()
	return block
}

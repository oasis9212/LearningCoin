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
	block []block
}

func main() {
	// go언어의 string 값은 불변하다.

	genesisBlcok := block{"Genesis Block", "", ""}
	//genesisBlcok.hash=  hashfunction(genesisBlcok.data+genesisBlcok.prevHash)   // SHA-256 해쉬 함수를 활용한다.
	hash := sha256.Sum256([]byte(genesisBlcok.data + genesisBlcok.prevHash)) //  SHA256 이미 구현되어있다.
	hexHash := fmt.Sprintf("%x", hash)                                       // string 값으로 변환한다.
	genesisBlcok.hash = hexHash                                              // 제네시스 블록 완성

	secondBlock := block{"Second Blocks", "", genesisBlcok.hash}
}

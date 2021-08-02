package utils

import (
	"bytes"
	"encoding/gob"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Tobytes(i interface{}) []byte {
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	HandleErr(encoder.Encode(i))
	return blockBuffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}

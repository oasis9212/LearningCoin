package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"ralo/utils"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			b1, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			b2, err := tx.CreateBucketIfNotExists([]byte(blocksBucket))

			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func Close() {
	DB().Close()
}

func SaveBlock(hash string, data []byte) {
	fmt.Printf("saving block %s \n Data %b", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dataBucket))
		err := b.Put([]byte("checkpoint"), data)
		return err
	})
	utils.HandleErr(err)
}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte("checkpoint"))
		return nil
	})
	return data
}

func Block(hash string) []byte {
	var data []byte

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}

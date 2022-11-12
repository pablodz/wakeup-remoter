package db

import (
	"log"

	"github.com/dgraph-io/badger"
)

var DATABASE *badger.DB = nil

func GetDatabase() *badger.DB {
	if DATABASE == nil {
		var err error = nil
		// Open the Badger database located in the /tmp/badger directory.
		// It will be created if it doesn't exist.
		DATABASE, err = badger.Open(badger.DefaultOptions("/tmp/badger"))
		if err != nil {
			log.Fatal(err)
		}
	}
	return DATABASE
}

func InsertKeyValue(key string, value string) {
	db := GetDatabase()
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteKey(key string) {
	db := GetDatabase()
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GetKeyValue(key string) string {
	db := GetDatabase()
	var item *badger.Item
	var err error
	err = db.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte(key))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	var valCopy []byte
	err = item.Value(func(val []byte) error {
		valCopy = append([]byte{}, val...)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return string(valCopy)
}

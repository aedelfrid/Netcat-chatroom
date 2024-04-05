package main

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
	"time"
)

type Message struct {
	username string
	timeStamp time.Time
	body string
}

type DB struct {
	*badger.DB
	
}

func initDB() *DB{
	d, err := badger.Open(badger.DefaultOptions("/data"))
	if err != nil {
		log.Fatal(err)
	}

	db := &DB{
		d,
	}

	return db
}

func (db *DB) getMessage() {
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("message"))
		
		return err
	})
}

func (db *DB) getUser() {
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("user"))

		return err
	})
}

func (db *DB) storeMessage() {
	err := db.Update(func(txn *badger.Txn) error {
  		// Your code here…
  		return err
	})
}

func (db *DB) storeUser() {
	err := db.Update(func(txn *badger.Txn) error {
  		// Your code here…
  		return err
	})
}


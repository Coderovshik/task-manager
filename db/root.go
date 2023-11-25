package db

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var taskDb *bolt.DB

func Open() {
	var err error
	taskDb, err = bolt.Open("task.db", 0644, &bolt.Options{ReadOnly: false})
	if err != nil {
		log.Fatal(err)
	}

	err = taskDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("list"))
		if err != nil {
			log.Fatal(err)
		}
		return err
	})
	if err != nil {
		fmt.Println("amogus")
		log.Fatal(err)
	}
}

func Close() {
	taskDb.Close()
}

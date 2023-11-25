package db

import (
	"encoding/binary"
	"log"

	"github.com/boltdb/bolt"
)

func AddTask(task string) (taskId int) {
	err := taskDb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("list"))

		id, _ := b.NextSequence()
		taskId = int(id)

		return b.Put(itob(taskId), []byte(task))
	})
	if err != nil {
		log.Fatal(err)
	}
	return
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

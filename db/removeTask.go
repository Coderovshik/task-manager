package db

import (
	"github.com/boltdb/bolt"
)

func RemoveTask(taskId string) {
	taskDb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("list"))

		b.Delete([]byte(taskId))

		return nil
	})
}

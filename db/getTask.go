package db

import "github.com/boltdb/bolt"

func GetTask(taskId string) string {
	var task string

	taskDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("list"))
		task = string(b.Get([]byte(taskId)))

		return nil
	})

	return task
}

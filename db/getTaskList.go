package db

import (
	"github.com/boltdb/bolt"
)

func GetTaskList() (taskList map[int][]string) {
	taskList = make(map[int][]string)
	taskDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("list"))

		c := b.Cursor()

		i := 0
		for k, v := c.First(); k != nil; k, v = c.Next() {
			i++
			id := string(k)
			taskList[i] = []string{id, string(v)}
		}

		return nil
	})

	return
}

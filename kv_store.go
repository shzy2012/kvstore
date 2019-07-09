package kvstore

/*
https://github.com/dgraph-io/badger
*/

import (
	"log"

	badger "github.com/dgraph-io/badger"
)

//KVStore 结构体
type KVStore struct {
	db   *badger.DB
	name string
}

//NewKVStore 创建KVStore
func NewKVStore(dbname string) *KVStore {
	db, err := badger.Open(badger.DefaultOptions(dbname))
	if err != nil {
		log.Fatal(err)
	}
	return &KVStore{
		db:   db,
		name: dbname,
	}
}

//Put 添加KV
func (k *KVStore) Put(key, value string) {

	err := k.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})

	if err != nil {
		log.Printf("%s\n", err)
	}
}

//Get 根据K获取V
func (k *KVStore) Get(key string) (string, error) {
	var result string
	
	err := k.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		item.Value(func(val []byte) error {
			result = string(val)
			return nil
		})

		return nil
	})
	return result, err
}

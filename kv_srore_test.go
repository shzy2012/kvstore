package kvstore

import (
	"fmt"
	"testing"
)

func Test_Run(t *testing.T) {
	db := NewKVStore("store")
	db.Put("key", "value")
	v, err := db.Get("key2")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println("[]=>", v)
}

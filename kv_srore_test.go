package kvstore

import (
	"fmt"
	"testing"
)

func Test_Run(t *testing.T) {
	db := NewKVStore("store")
	db.Put("key", "value")
	fmt.Println(db.Get("key"))
}

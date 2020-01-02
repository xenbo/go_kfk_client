package main

import (
	"fmt"
	rdkfk "github.com/xenbo/go_kfk_client/rdkfk"
	"strconv"
)

func main() {

	rdkfk.Cgo_init("192.168.1.1")

	db := rdkfk.Cgo_createdb("testdb", 100)

	var key string = "key_"
	var val string = "val_asdfasdfdsafasdfasdfasdfasdfasikodkfjklashdfjkasdhfk" +
		"jashdfjhasdjkfhajksdhfjkasdhfjkasdhfjkdashjfhasdjkfhasjkdfjaksd" +
		"hfjkasdhfjkasdhfjkasdfjkhasdjkasdfhjkasdhfjkasd"

	for i := 0; i < 100; i++ {
		var val1 string = val + strconv.Itoa(i)
		var key1 string = key + strconv.Itoa(i)
		rdkfk.Cgo_setkey(db, key1, val1)
	}

	for i := 0; i < 100; i++ {
		var key string = "key_"
		var key1 string = key + strconv.Itoa(i)

		v := rdkfk.Cgo_getkey(db, key1)
		fmt.Println(v)
	}
}

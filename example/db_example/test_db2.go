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
	var val string = "123123"

	rdkfk.Cgo_setkey(db, key, val)
	v := rdkfk.Cgo_getkey(db, key)

	n ,_:= strconv.Atoi(v)
	fmt.Println(n)
}

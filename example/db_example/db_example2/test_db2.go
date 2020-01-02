package main

import (
	"fmt"
	rdkfk "github.com/xenbo/go_kfk_client/rdkfk"
	"strconv"
)

func main() {

	rdkfk.CgoInit("192.168.1.1")

	db := rdkfk.OperateDb{}
	db.CreateDb("testdb", 100)

	var key string = "key_"
	var val string = "123123"

	db.SetKey( key, val)
	v := db.GetKey(key)

	n ,_:= strconv.Atoi(v)
	fmt.Println(n)
}

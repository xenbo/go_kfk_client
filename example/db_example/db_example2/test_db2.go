package main

import (
	"fmt"
	rdkfk "github.com/xenbo/go_kfk_client/rdkfk"
	"strconv"
)

func main() {

	glc:= rdkfk.GlobeCleaner{}
	glc.SetKafkaAddr("192.168.1.1")
	glc.Init()

	db := rdkfk.OperateDb{}
	db.CreateDb("testdb", 100)

	var key string = "key_"
	var val string = "123123"

	db.SetKey( key, val)
	v := db.GetKey(key)

	n ,_:= strconv.Atoi(v)
	fmt.Println(n)
}

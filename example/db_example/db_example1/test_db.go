package main

import (
	"fmt"
	rdkfk "github.com/xenbo/go_kfk_client/rdkfk"
	"strconv"
	"time"
)

func main() {

	glc:= rdkfk.GlobeCleaner{}
	glc.SetKafkaAddr("192.168.1.1")
	glc.Init()

	db := rdkfk.OperateDb{}
	db.CreateDb("testdb", 100)

	var key string = "key_"
	var val string = "val_asdfasdfdsafasdfasdfasdfasdfasikodkfjklashdfjkasdhfk" +
		"jashdfjhasdjkfhajksdhfjkasdhfjkasdhfjkdashjfhasdjkfhasjkdfjaksd" +
		"hfjkasdhfjkasdhfjkasdfjkhasdjkasdfhjkasdhfjkasd"

	for i := 0; i < 100; i++ {
		var val1 string = val + strconv.Itoa(i)
		var key1 string = key + strconv.Itoa(i)
		db.SetKey(key1, val1)
	}

	for i := 0; i < 100; i++ {
		var key string = "key_"
		var key1 string = key + strconv.Itoa(i)

		v := db.GetKey(key1)
		fmt.Println(v)
	}

	time.Sleep(10* time.Second)
}

package main

import (
	"fmt"
	rdkfk "github.com/xenbo/go_kfk_client/rdkfk"
	"time"
)

func callbak(offset int64, topic string, msg string) {
	fmt.Println("callbak")
	fmt.Println(offset)
	fmt.Println(topic)
	fmt.Println(msg)
}



func main() {
	rdkfk.Cgo_init("192.168.1.172")

	c := rdkfk.Cgo_NewConsumer("192.168.1.172",callbak)
	rdkfk.Cgo_add_consume_topic("test", 200000, c)
	rdkfk.Cgo_add_consume_topic("test_topic", 200000, c)

	go rdkfk.Cgo_start_consumer(c)






	time.Sleep(10000 * time.Second)
}

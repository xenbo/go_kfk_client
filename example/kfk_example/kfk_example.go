package main

import (
	"fmt"
	rdkfk "github.com/xenbo/go_kfk_client/rdkfk"
	"strconv"
	"time"
)

var db *rdkfk.Cgo_db

func callbak(offset int64, topic string, msg string) {
	fmt.Println("callbak")
	fmt.Println(offset)
	fmt.Println(topic)
	fmt.Println(msg)

	rdkfk.Cgo_setkey(db, "offset"+topic, strconv.FormatInt(offset, 10))
}

func callbak1(offset int64, topic string, msg string) {
	fmt.Println("callbak111111111111")
	fmt.Println(offset)
	fmt.Println(topic)
	fmt.Println(msg)

	rdkfk.Cgo_setkey(db, "offset"+topic, strconv.FormatInt(offset, 10))
}

func main() {
	rdkfk.Cgo_init("192.168.1.172")
	db = rdkfk.Cgo_createdb("testdb", 100)

	c := rdkfk.Cgo_NewConsumer("192.168.1.172", callbak)
	rdkfk.Cgo_add_consume_topic("test_topicxxx", 0, c)
	rdkfk.Cgo_add_consume_topic("test_topicxxx1", 0, c)
	rdkfk.Cgo_add_consume_topic("test_topicxxx2", 0, c)
	go rdkfk.Cgo_start_consumer(c)

	c1 := rdkfk.Cgo_NewConsumer("192.168.1.172", callbak1)
	rdkfk.Cgo_add_consume_topic("test_topicxxx", 0, c1)
	rdkfk.Cgo_add_consume_topic("test_topicxxx1", 0, c1)
	rdkfk.Cgo_add_consume_topic("test_topicxxx2", 0, c1)
	go rdkfk.Cgo_start_consumer(c1)


	p := rdkfk.Cgo_NewProducer("192.168.1.172:9092")
	rdkfk.Cgo_add_produce_topic("test_topicxxx", p)
	rdkfk.Cgo_add_produce_topic("test_topicxxx1", p)
	rdkfk.Cgo_add_produce_topic("test_topicxxx2", p)

	for i := 1; i < 1000000; i++ {
		rdkfk.Cgo_send_msg("test_topicxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxhasdkfhaskjdfajskdfjkasdhfjkasdhfkhasdjkfhadsjkfhjkasfhjksadhfjksadhfjksadhfjkasdhfjhasdjkhasjdkfasdjkfhas", p)
		rdkfk.Cgo_send_msg("test_topicxxx1", "xxxxxxxxxxxxxxxxxxxxxxxxxhasdkfhaskjdfajskdfjkasdhfjkasdhfkhasdjkfhadsjkfhjkasfhjksadhfjksadhfjksadhfjkasdhfjhasdjkhasjdkfasdjkfhas", p)
		rdkfk.Cgo_send_msg("test_topicxxx2", "xxxxxxxxxxxxxxxxxxxxxxxxxhasdkfhaskjdfajskdfjkasdhfjkasdhfkhasdjkfhadsjkfhjkasfhjksadhfjksadhfjksadhfjkasdhfjhasdjkhasjdkfasdjkfhas", p)
	}

	time.Sleep(10000 * time.Second)
}

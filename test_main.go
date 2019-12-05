package main

import (
	"rdkfk"
	"time"
)

//func main() {
//	c := rdkfk.Cgo_NewConsumer("127.0.0.1")
//	rdkfk.Cgo_add_consume_topic("test_topicxxx", c)
//	go rdkfk.Cgo_start_consumer(c)
//
//	time.Sleep(2 * time.Second)
//
//	p := rdkfk.Cgo_NewProducer("127.0.0.1")
//	rdkfk.Cgo_add_produce_topic("test_topicxxx", p)
//	for i := 1; i < 1000000; i++ {
//		rdkfk.Cgo_send_msg("test_topicxxx", "xxxxxxxxxxxxx", p)
//	}
//	rdkfk.Cgo_flush(p)
//
//	time.Sleep(1000 * time.Second)
//}


func main() {
	c := rdkfk.Cgo_NewConsumer("127.0.0.1")
	rdkfk.Cgo_add_consume_topic("test_topicxxx", c)
	go rdkfk.Cgo_start_consumer(c)

	time.Sleep(2 * time.Second)

	p := rdkfk.Cgo_NewProducer("127.0.0.1")
	rdkfk.Cgo_add_produce_topic("test_topicxxx", p)
	for i := 1; i < 1000000; i++ {
		rdkfk.Cgo_send_msg("test_topicxxx", "xxxxxxxxxxxxx", p)
	}
	rdkfk.Cgo_flush(p)

	time.Sleep(1000 * time.Second)
}

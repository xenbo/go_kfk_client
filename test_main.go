package main

import (
	rdkfk "./rdkfk"
	"time"
)

func main() {
	rdkfk.Cgo_init()

	c := rdkfk.Cgo_NewConsumer("192.168.1.172")
	rdkfk.Cgo_add_consume_topic("test_topicxxx", 10000000, c)
	rdkfk.Cgo_add_consume_topic("test_topicxxx1", 10000000, c)
	rdkfk.Cgo_add_consume_topic("test_topicxxx2", 10000000, c)
	go rdkfk.Cgo_start_consumer(c)

	time.Sleep(2 * time.Second)

	p := rdkfk.Cgo_NewProducer("192.168.1.172")
	rdkfk.Cgo_add_produce_topic("test_topicxxx", p)
	rdkfk.Cgo_add_produce_topic("test_topicxxx1", p)
	rdkfk.Cgo_add_produce_topic("test_topicxxx2", p)
	for i := 1; i < 1000000; i++ {
		rdkfk.Cgo_send_msg("test_topicxxx", "xxxxxxxxxxxxx", p)
		rdkfk.Cgo_send_msg("test_topicxxx1", "xxxxxxxxxxxxx", p)
		rdkfk.Cgo_send_msg("test_topicxxx2", "xxxxxxxxxxxxx", p)

		if i%1000 == 0 {
			rdkfk.Cgo_flush(p)
		}
	}

	time.Sleep(10000 * time.Second)
}

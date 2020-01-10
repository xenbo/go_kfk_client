package main

import (
	"encoding/json"
	"fmt"
	"github.com/xenbo/go_kfk_client/rdkfk"
	"os"
	"strconv"
	"time"
)

func callback0(offset int64, topic string, msg string, db *rdkfk.OperateDb) {
	fmt.Println(offset)
	fmt.Println(topic)
	fmt.Println(msg)
}

func create_testfile(file string) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err == nil {
		f.Close()
		os.Remove(file)
	}
	f, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()

	m := rdkfk.KafkaMsg{}

	for i := 0; i < 2000; i++ {
		m.Topic = "topic"
		m.Time = time.Now().Unix()
		m.Msg = "sadfsadfasdfsdafasdfasdfasdfsad1231231231231231[[][]_" + strconv.Itoa(i)
		b, _ := json.Marshal(m)

		f.WriteString(string(b) + "\n")
	}
}

func main() {
	create_testfile("./msg.txt")

	glc := rdkfk.GlobeCleaner{}
	glc.SetKafkaAddr("192.168.1.172")
	glc.Init()


	kc2 := rdkfk.KafkaClient{}
	kc2.SetCallBack(callback0)
	kc2.NewConsumer()
	kc2.AddConsumeTopic("topic", 0)
	go kc2.StartConsumer()


	kc3 := rdkfk.KafkaClient{}
	kc3.NewProducer()

	rdkfk.ProduceFile(kc3, "./msg.txt")

	time.Sleep(time.Second * 11130)
}

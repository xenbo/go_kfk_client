package main

import (
	"fmt"
	rdkfk "github.com/xenbo/go_kfk_client/rdkfk"
	"strconv"
	"time"
)

func callback0(offset int64, topic string, msg string, db *rdkfk.OperateDb) {
	fmt.Println("callback0000000000")
	fmt.Println(offset)
	fmt.Println(topic)
	fmt.Println(msg)

	db.SetKey("offset"+topic, strconv.FormatInt(offset, 10))
}



func main() {
	glc := rdkfk.GlobeCleaner{}
	glc.SetKafkaAddr("192.168.1.172")
	glc.Init()

	db := rdkfk.OperateDb{}
	db.CreateDb("offsetDB", 200)

	kc := rdkfk.KafkaClient{}
	kc.SetOpDb(&db)
	kc.SetCallBack(callback0)
	kc.NewConsumer()
	kc.AddConsumeTopic("goods_info_test", 0)

	go kc.StartConsumer()



	time.Sleep(10000 * time.Second)
}

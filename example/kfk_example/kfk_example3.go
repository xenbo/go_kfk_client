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

func callback1(offset int64, topic string, msg string, db *rdkfk.OperateDb) {
	fmt.Println("callback111111111111")
	fmt.Println(offset)
	fmt.Println(topic)
	fmt.Println(msg)

	db.SetKey("offset"+topic, strconv.FormatInt(offset, 10))
}

func main() {
	glc := rdkfk.GlobeCleaner{}
	glc.SetKafkaAddr("192.168.1.146")
	glc.Init()

	db := rdkfk.OperateDb{}
	db.CreateDb("offsetDB", 200)

	//kc := rdkfk.KafkaClient{}
	//kc.SetOpDb(&db)
	//kc.SetCallBack(callback0)
	//kc.NewConsumer()
	//kc.AddConsumeTopic("test_topicxxx", 0)
	//kc.AddConsumeTopic("test_topicxxx1", 0)
	//kc.AddConsumeTopic("test_topicxxx2", 0)
	//go kc.StartConsumer()

	//kc2 := rdkfk.KafkaClient{}
	//kc2.SetOpDb(&db)
	//kc2.SetCallBack(callback1)
	//kc2.NewConsumer()
	//kc2.AddConsumeTopic("test_topicxxx", 0)
	//kc2.AddConsumeTopic("test_topicxxx1", 0)
	//kc2.AddConsumeTopic("test_topicxxx2", 0)
	//go kc2.StartConsumer()

	kc3 := rdkfk.KafkaClient{}
	kc3.NewProducer()
	kc3.AddProduceTopic("test_topicxxx")
	kc3.AddProduceTopic("test_topicxxx1")
	kc3.AddProduceTopic("test_topicxxx2")

	for i := 1; i < 3; i++ {
		kc3.SendMsgWithCache("test_topicxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxhasdkfhaskjdfajskdfjkasdhfjkasdhfkhasdjkfhadsjkfhjkasfhjksadhfjksadhfjksadhfjkasdhfjhasdjkhasjdkfasdjkfhas")
		kc3.SendMsgWithCache("test_topicxxx1", "xxxxxxxxxxxxxxxxxxxxxxxxxhasdkfhaskjdfajskdfjkasdhfjkasdhfkhasdjkfhadsjkfhjkasfhjksadhfjksadhfjksadhfjkasdhfjhasdjkhasjdkfasdjkfhas")
		kc3.SendMsgWithCache("test_topicxxx2", "xxxxxxxxxxxxxxxxxxxxxxxxxhasdkfhaskjdfajskdfjkasdhfjkasdhfkhasdjkfhadsjkfhjkasfhjksadhfjksadhfjksadhfjkasdhfjhasdjkhasjdkfasdjkfhas")
	}

	time.Sleep(10000 * time.Second)
}

package rdkfk

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type KafkaMsg struct {
	Topic  string `json:"topic"`
	Msg    string `json:"msg"`
	Time   int64  `json:"time"`
	Offset int64  `json:"offset"`
}

func ProduceFile(kc KafkaClient, filename string) int {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n := 0
	rd := bufio.NewReader(f)
	for {
		lineMsg, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		var km KafkaMsg
		json.Unmarshal([]byte(lineMsg), &km)
		fmt.Println(km.Topic, km.Msg)

		kc.AddProduceTopic(km.Topic)
		kc.SendMsgWithCache(km.Topic, km.Msg)
		n++
	}

	return n
}

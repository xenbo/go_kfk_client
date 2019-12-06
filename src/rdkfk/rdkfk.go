package rdkfk

/*
#cgo CXXFLAGS: -std=c++14
#cgo LDFLAGS: -L.  -lkfclient -lrdkafka -lssl -lcrypto -ldl -lm -lz -lstdc++ -lpthread
#include "for_go.h"
*/
import "C"
import (
	"fmt"
)

type Cgo_Producer C.Producer_t
type Cgo_Consumer C.Consumer_t


func Cgo_NewProducer(addr string) *Cgo_Producer {
	p := C.create_producer(C.CString(addr))
	return (*Cgo_Producer)(p)
}

func Cgo_add_produce_topic(topic string, p *Cgo_Producer) {
	C.add_produce_topic((*C.Producer_t)(p), C.CString(topic))
}

func Cgo_send_msg(topic string, msg string, p *Cgo_Producer) {
	C.send_msg((*C.Producer_t)(p), C.CString(msg), C.CString(topic))
}

func Cgo_flush(p *Cgo_Producer) {
	C.flush((*C.Producer_t)(p))
}

func Cgo_NewConsumer(addr string) *Cgo_Consumer {
	p := C.create_consumer(C.CString(addr))
	return (*Cgo_Consumer)(p)
}

func Cgo_add_consume_topic(topic string, offset int64, c *Cgo_Consumer) {
	C.add_consume_topic((*C.Consumer_t)(c), C.CString(topic) , C.longlong(offset))
}

func Cgo_start_consumer(c *Cgo_Consumer) {
	C.start_consumer((*C.Consumer_t)(c))
}

//export Cgo_comsumer_callback
func Cgo_comsumer_callback(topic *C.char, offset C.longlong, msg *C.char, len C.int) {
	fmt.Println(C.GoString(topic))
	fmt.Println(offset)
	fmt.Println(C.GoStringN(msg, len))
	//fmt.Println(len)
}

func Cgo_init()  {
	C.init()
}


//type Cgo_db C.Storage_t
//
//func Cgo_createdb(dbnanme string, mmsize int) *Cgo_db {
//	db := C.storage_init(C.CString(dbnanme), C.int(mmsize))
//	return (*Cgo_db)(db)
//}
//
//func Cgo_setkey(db *Cgo_db, key string, val string) {
//	C.storage_setkey((*C.Storage_t)(db), C.CString(key), C.int(len(key)), C.CString(val), C.int(len(val)))
//}
//
//func Cgo_getkey(db *Cgo_db, key string) string {
//	var bytes = make([]byte, 200)
//	C.storage_getkey((*C.Storage_t)(db), C.CString(key), C.int(len(key)), unsafe.Pointer(&bytes[0]), C.int(len(bytes)))
//	val := string(bytes)
//	return val
//}
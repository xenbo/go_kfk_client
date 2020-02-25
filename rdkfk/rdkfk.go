package rdkfk

/*
#cgo CXXFLAGS: -std=c++17
#cgo linux  LDFLAGS:  -L ../lib -lsyrdkafka  -ldl -lm -lz -lstdc++ -lstdc++fs -lpthread
#cgo darwin LDFLAGS:  -L ../lib -lsyrdkafka  -ldl -lm -lz -lc++ -lpthread
#include "for_go.h"
*/
import "C"
import "sync"

type CgoBackFunc func(offset int64, topic string, msg string, db* OperateDb)

type CgoProducer C.Producer_t
type CgoConsumer C.Consumer_t

var OnConComeMsg = make(map[uint64] *KafkaClient)
var fucLock sync.Mutex

func CgoSetKafkaAddr(addr string)  {
	C.set_kfk_addr(C.CString(addr))
}

func CgoRecover(){
	C.recover()
}

func CgoNewProducer() *CgoProducer {
	p := C.create_producer()
	return (*CgoProducer)(p)
}

func CgoAddProduceTopic(topic string, p *CgoProducer) {
	C.add_produce_topic((*C.Producer_t)(p), C.CString(topic))
}

func CgoSendMsg(topic string, msg string, p *CgoProducer) {
	C.send_msg((*C.Producer_t)(p), C.CString(msg), C.CString(topic))
}
func CgoSendMsgWithCache(topic string, msg string, p *CgoProducer){
	C.send_msg_with_cache((*C.Producer_t)(p), C.CString(msg), C.CString(topic))
}

func CgoFlush(p *CgoProducer) {
	C.flush((*C.Producer_t)(p))
}

func CgoNewConsumer(kc * KafkaClient) *CgoConsumer {
	p := C.create_consumer()
	fucLock.Lock()
	OnConComeMsg[CgoGetConsumerHashCode((*CgoConsumer)(p))] = kc
	fucLock.Unlock()
	return (*CgoConsumer)(p)
}

func CgoGetConsumerHashCode(c *CgoConsumer) uint64 {
	return uint64(C.get_consumer_hash_code((*C.Consumer_t)(c)))
}

func CgoAddConsumeTopic(topic string, offset int64, c *CgoConsumer) {
	C.add_consume_topic((*C.Consumer_t)(c), C.CString(topic), C.longlong(offset))
}

func CgoStartConsumer(c *CgoConsumer) {
	C.start_consumer((*C.Consumer_t)(c))
}

//export CgoConsumerCallback
func CgoConsumerCallback(topic *C.char, offset C.longlong, msg *C.char, len C.int, consumer_hashcode C.ulonglong) {
	_topic := C.GoString(topic)
	_offset := int64(offset)
	_msg := C.GoStringN(msg, len)
	_hashcode := uint64(consumer_hashcode)

	fucLock.Lock()
	kc, ok := OnConComeMsg[_hashcode]
	fucLock.Unlock()
	if ok {
		kc.fuc(_offset, _topic, _msg, kc.db)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////

type GlobeCleaner struct {
	kafkaAddr string
}

func (gcl *GlobeCleaner )SetKafkaAddr(addr string){
	gcl.kafkaAddr = addr
}

func (gcl GlobeCleaner) Init()  {
	CgoSetKafkaAddr(gcl.kafkaAddr)
	CgoRecover()
}

type KafkaClient struct {
	p *CgoProducer
	c *CgoConsumer
	db *OperateDb
	fuc CgoBackFunc
}

func (kc *KafkaClient) NewProducer() *KafkaClient {
	kc.p = CgoNewProducer()
	return kc
}

func (kc KafkaClient) AddProduceTopic(topic string) {
	CgoAddProduceTopic(topic, kc.p)
}

func (kc KafkaClient) SendMsg(topic string, msg string) {
	CgoSendMsg(topic, msg, kc.p)
}

func (kc KafkaClient)SendMsgWithCache(topic string, msg string) {
	CgoSendMsgWithCache(topic, msg, kc.p)
}

func (kc KafkaClient) KFlush() {
	CgoFlush(kc.p)
}

func (kc *KafkaClient) NewConsumer() *KafkaClient {
	kc.c = CgoNewConsumer(kc)
	return kc
}

func (kc *KafkaClient) SetCallBack(fuc_ CgoBackFunc) ()  {
	kc.fuc = fuc_
}

func (kc KafkaClient) GetConsumerHashCode() uint64 {
	return CgoGetConsumerHashCode(kc.c)
}

func (kc KafkaClient) AddConsumeTopic(topic string, offset int64) {
	CgoAddConsumeTopic(topic, offset, kc.c)
}

func (kc KafkaClient) StartConsumer() {
	CgoStartConsumer(kc.c)
}

func (kc *KafkaClient) SetOpDb(db *OperateDb)  {
	kc.db = db
}

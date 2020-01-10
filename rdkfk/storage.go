package rdkfk

/*
#cgo CXXFLAGS: -std=c++17
#cgo LDFLAGS: -L ../lib -lsyrdkafka -lssl -lcrypto -ldl -lm -lz -lstdc++ -lstdc++fs -lpthread
#include "for_go.h"
*/
import "C"
import "unsafe"

type CgoDb C.Storage_t

func CgoCreateDb(dbname string, mmSize int) *CgoDb {
	db := C.storage_init(C.CString(dbname), C.int(mmSize))
	return (*CgoDb)(db)
}

func CgoSetKey(db *CgoDb, key string, val string) {
	C.storage_setkey((*C.Storage_t)(db), C.CString(key), C.int(len(key)), C.CString(val), C.int(len(val)))
}

func CgoGetKey(db *CgoDb, key string) string {
	var bytes = make([]byte, 1024)
	len := C.storage_getkey((*C.Storage_t)(db), C.CString(key), C.int(len(key)), unsafe.Pointer(&bytes[0]), C.int(len(bytes)))
	if len > 0 {
		return string(bytes[0:int32(len)])
	}
	return ""
}

type OperateDb struct {
	db *CgoDb
}

func (odb *OperateDb) CreateDb(dbname string, mmSize int) *OperateDb {
	odb.db = CgoCreateDb(dbname, mmSize)
	return odb
}

func (odb OperateDb) SetKey(key string, val string) {
	CgoSetKey(odb.db, key, val)
}

func (odb OperateDb) GetKey(key string) string {
	return CgoGetKey(odb.db, key)
}

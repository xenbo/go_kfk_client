package rdkfk

/*
#cgo CXXFLAGS: -std=c++14
#cgo LDFLAGS: -L../lib -lrdkafka -lssl -lcrypto -ldl -lm -lz -lstdc++ -lpthread
#include "for_go.h"
*/
import "C"
import "unsafe"

type Cgo_db C.Storage_t

func Cgo_createdb(dbnanme string, mmsize int) *Cgo_db {
	db := C.storage_init(C.CString(dbnanme), C.int(mmsize))
	return (*Cgo_db)(db)
}

func Cgo_setkey(db *Cgo_db, key string, val string) {
	C.storage_setkey((*C.Storage_t)(db), C.CString(key), C.int(len(key)), C.CString(val), C.int(len(val)))
}

func Cgo_getkey(db *Cgo_db, key string) string {
	var bytes = make([]byte, 1024)
	C.storage_getkey((*C.Storage_t)(db), C.CString(key), C.int(len(key)), unsafe.Pointer(&bytes[0]), C.int(len(bytes)))
	val := (*string)(unsafe.Pointer(&bytes))
	return *val
}

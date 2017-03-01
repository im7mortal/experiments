package library

//#include <stdlib.h>
//#cgo CPPFLAGS: -std=c++11
//#cgo CFLAGS: -w
//#cgo LDFLAGS: -lpthread
//#include "c_wrapper.h"
import "C"

import (
	"unsafe"
)

func Init() int {
	return int(C.initC())
}

func Get(a float32) {
	C.setC(C.int(a))
}

func Sent() {
	lol := []C.int{1, 56, 99}
	l := ([]C.int)(lol)
	C.sentC((*C.int)(unsafe.Pointer(&l[0])), C.int(len(lol)), C.int(unsafe.Sizeof(int8(0))))
}

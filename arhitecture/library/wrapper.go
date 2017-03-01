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

func Sent(a float32) {
	lol := []C.int{C.int(int(a)), 56, 99}
	l := ([]C.int)(lol)
	C.sentC((*C.int)(unsafe.Pointer(&l[0])), C.int(len(lol)), C.int(a))
}

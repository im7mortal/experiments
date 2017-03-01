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

type Context * C.int

func StartCalc(ctx Context) int {

	return int(C.initC(ctx))
}

func GetContext() Context {
	return C.GetContextC()
}

func SentDeffer(a float32) {
	lol := []C.int{C.int(int(a)), 56, 99}
	l := ([]C.int)(lol)
	C.sentC((*C.int)(unsafe.Pointer(&l[0])), C.int(len(lol)), C.int(a))
}
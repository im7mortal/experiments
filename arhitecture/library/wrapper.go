package library

//#include <stdlib.h>
//#cgo CPPFLAGS: -std=c++11
//#cgo CFLAGS: -w
//#cgo LDFLAGS: -lpthread
//#include "c_wrapper.h"
import "C"

func Init() int {
	return int(C.initC())
}

func Get(a float32) {
	C.setC(C.int(a))
}

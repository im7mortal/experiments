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

func StartCalc(ctx Context, data []float32, n, fov, w, h int) int {
	return int(C.start_calculation_c(
		ctx,
		(*C.float)(unsafe.Pointer(&data[0])),
		C.int(n),
		C.int(fov),
		C.int(w),
		C.int(h),
	))
}

func GetContext() Context {
	return C.GetContextC()
}

func SentDeffer(ctx Context, a float32) {
	lol := []C.int{C.int(int(a)), 56, 99}
	l := ([]C.int)(lol)
	println("SEND GO DATA")
	C.sentC(ctx, (*C.int)(unsafe.Pointer(&l[0])), C.int(len(lol)), C.int(a))
}

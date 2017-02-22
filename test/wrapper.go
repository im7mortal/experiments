package cglogtest

/*
#include <stdlib.h>
#cgo CPPFLAGS: -std=c++11
#cgo CFLAGS: -w
#cgo LDFLAGS: -L/home/petr/go/src/github.com/sturfeeinc/cglog/test/ -lcglog
#include "c_wrapper.h"
*/
import "C"

func Test()  {
	C.Test()
}

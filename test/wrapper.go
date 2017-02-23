package cglogtest

/*
#include <stdlib.h>
#cgo CPPFLAGS: -std=c++11
#cgo CFLAGS: -w
#cgo LDFLAGS: -lcglog
#include "c_wrapper.h"
*/
import "C"

func Test()  {
	C.Test()
}

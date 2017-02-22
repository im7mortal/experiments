package main

import (
	"C"
	"github.com/golang/glog"
)

//export Info
func Info(s *C.char) {
	glog.Info(C.GoString(s))
}

//export Warning
func Warning(s *C.char) {
	glog.Warning(C.GoString(s))
}

//export Error
func Error(s *C.char) {
	glog.Error(C.GoString(s))
}

//export Fatal
func Fatal(s *C.char) {
	glog.Fatal(C.GoString(s))
}

func main() {}

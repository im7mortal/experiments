package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/sturfeeinc/cglog/test"
	"time"
)

func init() {
	flag.Parse()
	//flag.Lookup("alsologtostderr").Value.Set("true")
}

func main() {
	time.Sleep(2 * time.Second)
	glog.Info("Started from go")
	cglogtest.Test()
	glog.Info("Finished from go")
	glog.Info("Finished from go")
	time.Sleep(2 * time.Second)
}

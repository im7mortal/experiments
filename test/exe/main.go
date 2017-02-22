package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/sturfeeinc/cglog/test"
)

func main() {
	flag.Parse()
	glog.Info("Started from go")
	cglogtest.Test()
	glog.Info("Finished from go")
}

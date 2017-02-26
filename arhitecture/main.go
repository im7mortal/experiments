package main

import (
	"github.com/sturfeeinc/cglog/arhitecture/library"
)

func main() {
	success := library.Init()
	if success != 0 {
		panic("Initialization error!")
	}
	for z := float32(0.); z < 90; z += 3 {
		library.Get(z)
	}
}

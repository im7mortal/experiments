package main

import (
	"github.com/sturfeeinc/cglog/arhitecture/library"
	"time"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	success := library.Init()
	if success != 0 {
		panic("Initialization error!")
	}
	for z := float32(0.); z < 90; z += 3 {
		wg.Add(1)
		go func(z float32) {
			library.Get(z)
			wg.Done()
		} (z)
	}
	for z := float32(0.); z < 90; z += 3 {
		time.Sleep(time.Second * 4)
		library.Sent(z)
	}


	wg.Wait()
}

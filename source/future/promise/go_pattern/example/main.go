package main

import (
	promise "github.com/sturfeeinc/cglog/source/future/promise/go_pattern"
	"time"
)

func handler(p *promise.Promise)  {
	println("ask user info")
	u, ok := p.Get().(user)
	if ok {
		println("Hi from " + u.name)
	}
	println("got user info")
}

type user struct {
	name string
}

func main() {
	name := "Petr"
	var GetUserInfoFromDelayedSource  promise.DeferReturnFunc
	GetUserInfoFromDelayedSource = func(p *promise.Promise) {
		time.Sleep(time.Second)
		p.Done(user{name})
	}
	p := promise.New(GetUserInfoFromDelayedSource)

	go handler(p)
	go handler(p)
	go handler(p)

	time.Sleep(2 * time.Second)
}

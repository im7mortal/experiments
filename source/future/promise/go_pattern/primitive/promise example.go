package promise

import (
	"sync/atomic"
)

type Promise struct {
	value interface{}
	ret chan struct{}
	finished *int32
}

type DeferReturnFunc func(p *Promise)

func New(drf DeferReturnFunc) *Promise {
	p := new(Promise)
	p.ret = make(chan struct{})
	i := int32(0)
	p.finished = &i
	go drf(p)
	return p
}

func (p *Promise) Get() interface{} {
	finished := (atomic.LoadInt32(p.finished) != 0)
	if !finished {
		<-p.ret
	}
	return p.value
}

func (p *Promise) Done(i interface{}) {
	finished := (atomic.LoadInt32(p.finished) != 0)
	if finished {
		panic("Promise was returned already.")
	}
	p.value = i
	atomic.StoreInt32(p.finished, 1)
	close(p.ret)
}

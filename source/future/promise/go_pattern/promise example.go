package promise

import "sync"

type Promise struct {
	value interface{}
	wg    *sync.WaitGroup
}

type DeferReturnFunc func(p *Promise)

func New(drf DeferReturnFunc) *Promise {
	p := new(Promise)
	p.wg = new(sync.WaitGroup)
	p.wg.Add(1)
	go drf(p)
	return p
}

func (p *Promise) Get() interface{} {
	p.wg.Wait()
	return p.value
}

func (p *Promise) Done(i interface{}) {
	p.value = i
	p.wg.Done()
}

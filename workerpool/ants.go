package workerpool

import (
	"github.com/panjf2000/ants/v2"
)

type AntsWorkerPool struct {
	p          *ants.Pool
	errHandler func(err error)
}

func NewAntsWorkerPool(cap int, eh func(err error), ph PanicHandler) (*AntsWorkerPool, error) {
	if ph == nil {
		ph = DefaultPanicHandler
	}
	p, err := ants.NewPool(cap, ants.WithPanicHandler(ph))
	if err != nil {
		return nil, err
	}
	return &AntsWorkerPool{p: p, errHandler: eh}, nil
}

func (a *AntsWorkerPool) Submit(f func()) {
	if err := a.p.Submit(f); err != nil {
		a.errHandler(err)
	}
}

func (a *AntsWorkerPool) Count() int {
	return a.p.Running()
}

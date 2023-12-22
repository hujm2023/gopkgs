package workerpool

import (
	"context"

	"github.com/bytedance/gopkg/util/gopool"
)

type GoPool struct {
	p gopool.Pool
}

func NewGoPool(name string, cap int32, ph PanicHandler) *GoPool {
	if ph == nil {
		ph = DefaultPanicHandler
	}
	p := gopool.NewPool(name, cap, gopool.NewConfig())
	p.SetPanicHandler(func(_ context.Context, msg any) {
		ph(msg)
	})
	return &GoPool{p: p}
}

func (g *GoPool) Count() int {
	return int(g.p.WorkerCount())
}

func (g *GoPool) Submit(f func()) {
	g.p.Go(f)
}

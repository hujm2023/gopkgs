package workerpool

import (
	"github.com/hujm2023/hlog"
)

type WorkerPool interface {
	// Submit starts a task
	Submit(f func())

	// Count returns how many workers(goroutine) are running.
	Count() int
}

type PanicHandler func(msg any)

func DefaultPanicHandler(msg any) {
	hlog.Errorf("[WorkerPool] panic by %v", msg)
}

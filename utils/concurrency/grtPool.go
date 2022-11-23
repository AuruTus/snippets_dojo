package concurrency

import (
	"fmt"
)

type Pool interface {
	Start(task func()) error
}

type GrtPool struct {
	tasks     chan func()
	doneList  []chan struct{}
	size      int
	available Semaphore
}

func InitGrtPool(size, capacity int) (*GrtPool, error) {
	if size <= 0 || capacity <= 0 || size > capacity {
		return nil, fmt.Errorf("invalid pool size")
	}

	// no context needed
	sema, _ := NewDummySemaphore(int64(size))
	pool := GrtPool{
		tasks:     make(chan func(), size),
		doneList:  make([]chan struct{}, 0, 8),
		size:      size,
		available: sema,
	}

	for i := 0; i < size; i++ {
		done := make(chan struct{})
		pool.doneList = append(pool.doneList, done)
		go func() {
			for {
				select {
				case t := <-pool.tasks:
					t()
				case <-done:
					return
				}
			}
		}()
	}

	return &pool, nil
}

func (pl *GrtPool) emptySize() bool {
	return pl.size <= 0
}

func (pl *GrtPool) Start(task func()) (ok bool) {
	if pl == nil || pl.emptySize() {
		return false
	}
	ok = pl.available.TryAcquire(1)
	if !ok {
		return ok
	}
	pl.tasks <- task
	return ok
}

func (pl *GrtPool) Cancel() error {
	return nil
}

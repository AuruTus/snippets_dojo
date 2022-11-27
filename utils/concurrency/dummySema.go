package concurrency

import (
	"container/list"
	"context"
	"fmt"
	"sync"
)

type dummySema struct {
	size    int64
	cur     int64
	m       sync.Mutex
	waiters *list.List
}

var _ Semaphore = (*dummySema)(nil)

func NewDummySemaphore(n int64) (*dummySema, error) {
	if n < 0 {
		return nil, fmt.Errorf("non-negative integer %d passed in", n)
	}
	return &dummySema{
		size:    n,
		cur:     n,
		waiters: list.New(),
	}, nil
}

func (ds *dummySema) Acquire(ctx context.Context, n int64) error {
	return ds.acquire(n)
}

func (ds *dummySema) acquire(n int64) error {
	ds.m.Lock()
	// invalid n passed in
	if n > ds.size {
		ds.m.Unlock()
		return fmt.Errorf("token amount overflow")
	}
	// goroutine get n tokens and continue
	if n <= ds.size-ds.cur && ds.waiters.Len() == 0 {
		ds.cur += n
		ds.m.Unlock()
		return nil
	}

	// join goroutine into blocking queue
	ready := make(chan struct{})
	w := waiter{n: n, ready: ready}
	ds.waiters.PushBack(w)

	ds.m.Unlock()

	// goroutine blocks here
	<-ready
	return nil
}

func (ds *dummySema) TryAcquire(n int64) bool {
	ds.m.Lock()
	success := ds.cur+n <= ds.size && ds.waiters.Len() == 0
	if success {
		ds.cur += n
	}
	ds.m.Unlock()
	return success
}

func (ds *dummySema) Release(n int64) {
	ds.m.Lock()
	ds.cur -= n
	if ds.cur < 0 {
		ds.m.Unlock()
		panic("dummySema: released more than held")
	}
	ds.notifyWaiters()
	ds.m.Unlock()
}

func (ds *dummySema) notifyWaiters() {
	for {
		next := ds.waiters.Front()
		if next == nil {
			break
		}

		w := next.Value.(waiter)
		if w.n > ds.size-ds.cur {
			break
		}

		ds.cur += w.n
		ds.waiters.Remove(next)
		close(w.ready)
	}
}

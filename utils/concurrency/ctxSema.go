package concurrency

import (
	"container/list"
	"context"
	"fmt"
	"sync"
)

type ctxSema struct {
	size    int64
	cur     int64
	m       sync.Mutex
	waiters *list.List
}

var _ Semaphore = &ctxSema{}

func NewContextSemaphore(n int64) (*ctxSema, error) {
	if n < 0 {
		return nil, fmt.Errorf("non-negative integer %d passed in", n)
	}
	return &ctxSema{
		size:    n,
		cur:     n,
		waiters: list.New(),
	}, nil
}

func (cs *ctxSema) Acquire(ctx context.Context, n int64) error {
	cs.m.Lock()
	// invalid n passed in
	if n > cs.size {
		cs.m.Unlock()
		<-ctx.Done()
		return ctx.Err()
	}
	// goroutine get n tokens and continue
	if n <= cs.size-cs.cur && cs.waiters.Len() == 0 {
		cs.cur += n
		cs.m.Unlock()
		return nil
	}

	// join goroutine into blocking queue
	ready := make(chan struct{})
	w := waiter{n: n, ready: ready}
	elem := cs.waiters.PushBack(w)

	cs.m.Unlock()

	// goroutine blocks here
	select {
	case <-ctx.Done():
		err := ctx.Err()
		cs.m.Lock()
		select {
		case <-ready:
			err = nil
		default:
			isFront := cs.waiters.Front() == elem
			cs.waiters.Remove(elem)

			if isFront && cs.size > cs.cur {
				// ctx is cancelled, while there are still goroutines blocked in queue
				cs.notifyWaiters()
			}
		}
		cs.m.Unlock()
		return err
	case <-ready:
		return nil
	}
}

func (cs *ctxSema) TryAcquire(n int64) bool {
	cs.m.Lock()
	success := cs.cur+n <= cs.size && cs.waiters.Len() == 0
	if success {
		cs.cur += n
	}
	cs.m.Unlock()
	return success
}

func (cs *ctxSema) Release(n int64) {
	cs.m.Lock()
	cs.cur -= n
	if cs.cur < 0 {
		cs.m.Unlock()
		panic("ctxSema: released more than held")
	}
	cs.notifyWaiters()
	cs.m.Unlock()
}

func (cs *ctxSema) notifyWaiters() {
	for {
		next := cs.waiters.Front()
		if next == nil {
			break
		}

		w := next.Value.(waiter)
		if w.n > cs.size-cs.cur {
			break
		}

		cs.cur += w.n
		cs.waiters.Remove(next)
		close(w.ready)
	}
}

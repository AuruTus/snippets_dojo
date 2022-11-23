package concurrency

import "context"

type Semaphore interface {
	Acquire(ctx context.Context, n int64) error
	TryAcquire(n int64) bool
	Release(n int64)
	notifyWaiters()
}

type waiter struct {
	n     int64
	ready chan struct{}
}

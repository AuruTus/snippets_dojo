package concurrency

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type RWRequest int16

// One reader may get the read-right of another reader. May need some id number to control this
const (
	WRITE_REQUEST RWRequest = iota
	READ_REQUEST
)

type RWQueLock struct {
	rwRequests *list.List
	wQueue     chan struct{}
	rQueue     chan struct{}
	mu         sync.Mutex
	s          Semaphore
	writers    int
	readers    int
}

func NewRWQueLock(writers, readers int) (*RWQueLock, error) {
	if writers <= 0 || readers <= 0 {
		return nil, fmt.Errorf("invalid quantity of writer or reader")
	}
	s, _ := NewDummySemaphore(1)
	return &RWQueLock{
		rwRequests: list.New(),
		wQueue:     make(chan struct{}, writers),
		rQueue:     make(chan struct{}, readers),
		mu:         sync.Mutex{},
		s:          s,
		writers:    0,
		readers:    0,
	}, nil
}

func (l *RWQueLock) dispatch() {
	for l.rwRequests.Len() > 0 {
		l.mu.Lock()
		front := l.rwRequests.Front()
		l.rwRequests.Remove(front)
		l.mu.Unlock()

		// try not to dispatch the request until workers of the other type has finished.
		rq := front.Value.(RWRequest)
		for done := false; !done; {

			// refresh current workers' amounts
			l.mu.Lock()
			w := l.writers
			r := l.readers
			l.mu.Unlock()

			switch rq {
			case READ_REQUEST:
				// wait for the writer completing its work
				if w != 0 {
					time.Sleep(1 * time.Millisecond)
					break
				}
				l.rQueue <- struct{}{}
				done = true
			case WRITE_REQUEST:
				// wait for the readers or writer completing their work
				if r != 0 || w != 0 {
					time.Sleep(1 * time.Millisecond)
					break
				}
				l.wQueue <- struct{}{}
				done = true
			}
		}
	}
}

func (l *RWQueLock) tryDispatch(rqst RWRequest) {
	l.mu.Lock()
	l.rwRequests.PushBack(rqst)
	// increase worker amounts here
	switch rqst {
	case READ_REQUEST:
		l.readers++
	case WRITE_REQUEST:
		l.writers++
	}
	l.mu.Unlock()

	if !l.s.TryAcquire(1) {
		return
	}
	// get semaphore successfully
	l.dispatch()
	l.s.Release(1)
}

func (l *RWQueLock) ReadRequest() {
	go l.tryDispatch(READ_REQUEST)
	<-l.rQueue
}

func (l *RWQueLock) FinishRead() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.readers--
}

func (l *RWQueLock) WriteRequest() {
	go l.tryDispatch(WRITE_REQUEST)
	<-l.wQueue
}

func (l *RWQueLock) FinishWrite() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.writers--
}

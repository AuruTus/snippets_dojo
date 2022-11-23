package concurrency

import (
	"context"
	"fmt"
)

type RWRequest int16

// One reader may get the read-right of another reader. May need some id number to control this
const (
	WRITE_REQUEST RWRequest = iota
	READ_REQUEST
)

type RWQueLock struct {
	rwQueue chan RWRequest
	wQueue  <-chan RWRequest
	rQueue  <-chan RWRequest
	writers int
	readers int
}

func NewRWQueLock(writers, readers int) (*RWQueLock, error) {
	if writers <= 0 || readers <= 0 {
		return nil, fmt.Errorf("invalid quantity of writer or reader")
	}
	return &RWQueLock{
		rwQueue: make(chan RWRequest, (readers + writers)),
		wQueue:  make(<-chan RWRequest, writers),
		rQueue:  make(<-chan RWRequest, readers),
		writers: 0,
		readers: 0,
	}, nil
}

func (l *RWQueLock) ReadRequest(ctx context.Context) {
	go func() {
		l.rwQueue <- READ_REQUEST
	}()
	<-l.rQueue
}

func (l *RWQueLock) FinishRead() {
}

func (l *RWQueLock) WriteRequest(ctx context.Context) {
	l.rwQueue <- WRITE_REQUEST
	<-l.wQueue
}

func (l *RWQueLock) FinishWrite() {
}

package test_snippets

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	snippets "snippets_dojo/src"
	"syscall"
	"time"
)

type CloseChanTstr struct{}

var _ snippets.Tstr = (*CloseChanTstr)(nil)

func (tstr *CloseChanTstr) Test(ctx context.Context) error {
	const SIZE = 5

	os_signal := make(chan os.Signal)
	signal.Notify(os_signal, syscall.SIGINT, syscall.SIGTERM)
	input_chan := make(chan struct{}, SIZE)
	done := make(chan struct{})

	const (
		WAIT_DURATION_CNT       = 500
		CHILD_WAIT_DURATION_CNT = 1500
	)
	go func() {
		defer func() { done <- struct{}{} }()
		rcv_cnt := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("%e\n", ctx.Err())
				return
			case <-os_signal:
				return
			case _, ok := <-input_chan:
				if !ok {
					fmt.Printf("discard %d inputs; quit goroutine; \n", len(input_chan))
					return
				}
				rcv_cnt++
				fmt.Printf("child routine received! rcv_cnt: %d \n", rcv_cnt)
				time.Sleep(CHILD_WAIT_DURATION_CNT * time.Millisecond)
			}
			fmt.Printf("child waiting; buffured len: %d \n", len(input_chan))
		}
	}()

send_inputs_loop:
	for i := 0; i < SIZE; i++ {
		select {
		case <-ctx.Done():
			break send_inputs_loop
		case <-os_signal:
			break send_inputs_loop
		default:
			// fmt.Printf("send info\n")
			input_chan <- struct{}{}
		}
	}

	time.Sleep(WAIT_DURATION_CNT * time.Millisecond)
	fmt.Printf("complete sending inputs; start closing channel; \n")
	close(input_chan)
	fmt.Printf("complete closing channel; \n")
	<-done

	return nil
}

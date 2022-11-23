package snipets

import (
	"context"
	"snipets_dojo/utils/cfmt"
	"sync"
	"time"
)

type SelectOrderTstr struct{}

func (tstr *SelectOrderTstr) Test(ctx context.Context) error {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	closeChannels := func() {
		close(c3)
		close(c1)
		close(c2)
	}
	sendChannels := func() {
		c3 <- struct{}{}
		c1 <- struct{}{}
		c2 <- struct{}{}
	}
	channelEndingSwitcher := func(mode int32) {
		switch mode {
		case 1:
			closeChannels()
		case 2:
			sendChannels()
		default:
			closeChannels()
		}
	}

	m := sync.Mutex{}

	go func() {
		cfmt.Printf(ctx, "start cnt\n")
		// guranteen the main goroutine is blocked
		time.Sleep(2 * time.Second)
		channelEndingSwitcher(2)
		cfmt.Printf(ctx, "channel closed\n")
		m.Unlock()
	}()

	// await the channel signal
	m.Lock()
	cfmt.Printf(ctx, "enter channel\n")
	select {
	case <-c1:
		cfmt.Printf(ctx, "c1 done\n")
	case <-c2:
		cfmt.Printf(ctx, "c2 done\n")
	case <-c3:
		cfmt.Printf(ctx, "c3 done\n")
	}
	cfmt.Printf(ctx, "finish test\n")
	return nil
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	cfmt "snippets_dojo/utils/cfmt"
	ctxinfo "snippets_dojo/utils/ctx_info"
	"syscall"
	"time"
)

func main() {
	tstr := NewTesterEntry()

	/* NOTE: os signal channel's buffer is needed */
	os_signal := make(chan os.Signal, 1)
	signal.Notify(os_signal, syscall.SIGINT, syscall.SIGTERM)

	bg_ctx, bg_cancel := context.WithCancel(context.Background())
	defer bg_cancel()

	ctx, _ := ctxinfo.NewContextWithInfo(bg_ctx, ctxinfo.MAIN)

	tstr_done := make(chan struct{})

	// start snippet test
	start := time.Now()
	go func() {
		tstrCtx, _ := ctxinfo.NewContextWithInfo(ctx, ctxinfo.TESTER)
		err := tstr.Test(tstrCtx)
		if err != nil {
			err = fmt.Errorf("%T with error: %w", tstr, err)
			fmt.Println(err)
		}
		tstr_done <- struct{}{}
	}()

	// metrics
	cfmt.Printf(
		ctx,
		"%s starts at: %ss\n",
		cfmt.WarnStr("%T", tstr),
		cfmt.InfoStr("%d", start.Unix()),
	)

	/*
	 * Blocked for awaiting Tstr done or OS signal
	 */
	select {
	case <-os_signal:
	case <-tstr_done:
	}

	end := time.Now()
	elapsed := end.Sub(start)
	cfmt.Printf(
		ctx,
		"%s ends at %ss; elapse: %sns\n",
		cfmt.WarnStr("%T", tstr),
		cfmt.InfoStr("%d", end.Unix()),
		cfmt.InfoStr("%d", elapsed.Nanoseconds()),
	)
}

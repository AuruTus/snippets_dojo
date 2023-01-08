package test_snippets

import (
	"context"
	snippets "snippets_dojo/src"
	"snippets_dojo/utils/cfmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ChanRangeTstr struct{}

var _ snippets.Tstr = (*ChanRangeTstr)(nil)

var logger *cfmt.LogWrapper

func (tstr *ChanRangeTstr) Test(ctx context.Context) error {
	logger = cfmt.NewSimpleLogger(ctx)

	wg := sync.WaitGroup{}
	tasks := make(chan func(args ...any), 10)

	const GO_ROUTINE_NUM = 10
	for i := 0; i < GO_ROUTINE_NUM; i++ {
		wg.Add(1)
		grt_id, _ := uuid.NewUUID()
		go func() {
			defer wg.Done()
			time := time.NewTicker(100 * time.Millisecond)
			for elem := range tasks {
				t := <-time.C
				elem := elem
				elem(grt_id.ID(), t)
			}
		}()
	}

	const TASK_COUNTS = 100
	for i := 0; i < TASK_COUNTS; i++ {
		i := i
		tasks <- func(args ...any) {
			logger.Log(
				"grt %d | %s: I'm task %d",
				args[0].(uint32),
				args[1].(time.Time).String(),
				i,
			)
		}
	}
	/*
	 * NOTICE: If the tasks channel is not closed, the goroutine which
	 * is itrating the channel's task will blocked there until the main
	 * thread is ended (ctx ends or OS terminating signal)
	 */
	close(tasks)
	wg.Wait()

	return nil
}

package main

import (
	"context"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

func TestProducterConsumer(t *testing.T) {

	var (
		n      = 5
		empty  = semaphore.NewWeighted(int64(n))
		full   = semaphore.NewWeighted(int64(n))
		mutex  = semaphore.NewWeighted(1)
		buffer = make([]int, 0, n)
		ctx    = context.TODO()
		sets   = sync.Map{}
		cnt    = 0
	)

	full.Acquire(ctx, int64(n))

	producer := func(i int) {
		for {
			empty.Acquire(ctx, 1)
			mutex.Acquire(ctx, 1)
			buffer = append(buffer, i)
			mutex.Release(1)
			full.Release(1)
		}
	}

	consumer := func() {
		for {
			full.Acquire(ctx, 1)
			mutex.Acquire(ctx, 1)
			i := buffer[0]
			buffer = buffer[1:]
			cnt++
			mutex.Release(1)
			empty.Release(1)

			//use the item
			sets.Store(i, struct{}{})
		}
	}

	go producer(0)
	go consumer()
	time.Sleep(100 * time.Millisecond)

	if cnt < n {
		t.Errorf("expected %d, got %d", n, cnt)
	}
}

package monitor

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type ProducerConsumer struct {
	n       int
	count   int
	full    *semaphore.Weighted
	empty   *semaphore.Weighted
	buffers []int
	ctx     context.Context
}

func NewProducerConsumer(n int) *ProducerConsumer {
	p := &ProducerConsumer{
		n:       n,
		count:   0,
		full:    semaphore.NewWeighted(int64(1)),
		empty:   semaphore.NewWeighted(int64(1)),
		buffers: make([]int, 0, n),
		ctx:     context.TODO(),
	}

	p.full.Acquire(p.ctx, 1)
	p.empty.Acquire(p.ctx, 1)
	return p
}

func (pc *ProducerConsumer) Insert(i int) {
	if pc.count == pc.n {
		pc.full.Acquire(pc.ctx, 1)
	}

	pc.count++
	pc.buffers = append(pc.buffers, i)

	if pc.count == 1 {
		pc.empty.Release(1)
	}
}

func (pc *ProducerConsumer) Remove() int {
	var res int
	if pc.count == 0 {
		pc.empty.Acquire(pc.ctx, 1)
	}

	pc.count--
	res = pc.buffers[0]
	pc.buffers = pc.buffers[1:]

	if pc.count == pc.n-1 {
		pc.full.Release(1)
	}

	return res
}

package monitor

import (
	"context"
	"testing"
	"time"
)

func TestProducerConsumer(t *testing.T) {
	p := NewProducerConsumer(5)

	ctx := context.Background()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
			p.Insert(1)
		}
	}(ctx)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
			p.Remove()
		}
	}(ctx)

	time.Sleep(100 * time.Millisecond)
	ctx.Done()
}

package diningphilosophersproblem

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

func TestTwoChopstickSolution(t *testing.T) {
	var (
		mutex     = semaphore.NewWeighted(1)
		chopstick = []*semaphore.Weighted{
			semaphore.NewWeighted(1),
			semaphore.NewWeighted(1),
			semaphore.NewWeighted(1),
			semaphore.NewWeighted(1),
			semaphore.NewWeighted(1),
		}
		ctx = context.TODO()
	)

	p := func(i int) {
		mutex.Acquire(ctx, 1)
		chopstick[i].Acquire(ctx, 1)
		chopstick[(i+1)%5].Acquire(ctx, 1)
		mutex.Release(1)
		fmt.Println("Philosopher", i, "is eating")
		chopstick[i].Release(1)
		chopstick[(i+1)%5].Release(1)
		fmt.Println("Philosopher", i, "is thinking")
	}

	for i := 0; i < 5; i++ {
		go p(i)
	}

	// Wait for a while to let the goroutines finish
	time.Sleep(100 * time.Millisecond)
}

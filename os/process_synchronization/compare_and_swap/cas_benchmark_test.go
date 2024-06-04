package compareandswap

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Foo struct {
	v  int32
	mu sync.Mutex
}

func NewFoo() *Foo {
	return &Foo{}
}

func (foo *Foo) Inc() {
	foo.mu.Lock()
	foo.v++
	foo.mu.Unlock()
}

func (foo *Foo) IncWithCAS() {
	for {
		v := foo.v
		if atomic.CompareAndSwapInt32(&foo.v, v, v+1) {
			break
		}
	}
}

func BenchmarkIncWithLock(b *testing.B) {
	foo := NewFoo()
	wg := sync.WaitGroup{}

	n := b.N
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			foo.Inc()
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkIncWithCAS(b *testing.B) {

	foo := NewFoo()
	wg := sync.WaitGroup{}

	n := b.N
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			foo.IncWithCAS()
			wg.Done()
		}()
	}

	wg.Wait()
}

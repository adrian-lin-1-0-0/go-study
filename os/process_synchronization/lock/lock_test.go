package lock

import (
	"sync"
	"testing"
)

func NewNum() *Num {
	return &Num{
		locker: NewLocker(),
	}
}

type Num struct {
	v      int
	locker *Locker
}

func (num *Num) Inc() {
	num.locker.Acquire()
	num.v++
	num.locker.Release()
}

func Test_Lock(t *testing.T) {

	num := NewNum()
	wg := sync.WaitGroup{}

	n := 1000

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			num.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	if num.v != n {
		t.Errorf("num = %d; want %d", num.v, n)
	}
}

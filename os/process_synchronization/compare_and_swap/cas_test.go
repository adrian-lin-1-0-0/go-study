package compareandswap

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Locker struct {
	state *int32
}

const (
	nonLocked = 0
	locked    = 1
)

func NewLocker() *Locker {
	return &Locker{
		state: new(int32),
	}
}

func (l *Locker) Acquire() {
	for !CompareAndSwap(l.state, nonLocked, locked) {
	}
}

func (l *Locker) Release() {
	*l.state = nonLocked
}

func NewNum() *Num {
	return &Num{
		locker: NewLocker(),
	}
}

type Num struct {
	v      int32
	locker *Locker
}

func (num *Num) Inc() {
	num.locker.Acquire()
	num.v++
	num.locker.Release()
}

func (num *Num) IncWithCAS() {
	for {
		v := num.v
		if atomic.CompareAndSwapInt32(&num.v, v, v+1) {
			break
		}
	}
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

	if num.v != int32(n) {
		t.Errorf("num = %d; want %d", num.v, n)
	}
}

func Test_LockWithCAS(t *testing.T) {

	num := NewNum()
	wg := sync.WaitGroup{}

	n := 1000

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			num.IncWithCAS()
			wg.Done()
		}()
	}

	wg.Wait()

	if num.v != int32(n) {
		t.Errorf("num = %d; want %d", num.v, n)
	}
}

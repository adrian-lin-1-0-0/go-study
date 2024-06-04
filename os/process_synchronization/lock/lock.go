package lock

import (
	"sync/atomic"
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
	for !atomic.CompareAndSwapInt32(l.state, nonLocked, locked) {
	}
}

func (l *Locker) Release() {
	atomic.AddInt32(l.state, -locked)
}

# Mutex Lock

## Spinning Lock

用`golang`簡單實現一個自旋鎖:
```go
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
```


缺點是會忙等待，浪費CPU資源。

`golang`實際上如果無法`CAS`成功，會呼叫`lockSlow()`
> 進入`runtime`的`park`狀態，不會浪費CPU資源。

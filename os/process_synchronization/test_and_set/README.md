# TestAndSet

或稱`TestAndeSetLock`;是由硬體提供的原子指令，用來實現互斥鎖。

缺點:
- 不滿足讓權等待


實現:
> 很顯然,如果`lock`為`true`,FOR-LOOP會一直執行,直到`lock`為`false`為止。
```go
func TestAndSet(lock *bool) bool {
	old := *lock
	*lock = true
	return old
}
```

```go
lock := false

for TestAndSet(&lock) {
}
// critical section
lock = false                
```
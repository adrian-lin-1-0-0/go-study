# Swap

也叫`Exchange`或是`XCHG`，也是硬體支援的指令，邏輯跟`TestAndSet`一樣。

缺點:
- 不滿足讓權等待

```go
func Swap(a *bool, b *bool) {
	*a, *b = *b, *a
}
```

```go
lock, old := false, true

for old {
    Swap(&lock, &old)
}
// critical section
lock = false
```
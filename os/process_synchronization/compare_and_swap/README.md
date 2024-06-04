# Compare and swap (CAS)

```go
func CompareAndSwap(addr *int32, oldVal, newVal int32) bool {
	if *addr == oldVal {
		*addr = newVal
		return true
	}
	return false
}
```

### ABA問題
> ref: https://zh.wikipedia.org/zh-tw/%E6%AF%94%E8%BE%83%E5%B9%B6%E4%BA%A4%E6%8D%A2
ABA問題是無鎖結構實現中常見的一種問題，可基本表述為：

1. 進程P1讀取了一個數值A
2. P1被掛起(時間片耗盡、中斷等)，進程P2開始執行
3. P2修改數值A為數值B，然後又修改回A
4. P1被喚醒，比較後發現數值A沒有變化，程序繼續執行。

對於P1來說，數值A未發生過改變，但實際上A已經被變化過了，繼續使用可能會出現問題。在CAS操作中，由於比較的多是指針，這個問題將會變得更加嚴重。試想如下情況：
```
   top
    |
    V   
  0x0014
| Node A | --> |  Node X | --> ……
```
有一個棧(先入後出)中有top和節點A，節點A目前位於棧頂top指針指向A。現在有一個進程P1想要pop一個節點，因此按照如下無鎖操作進行
```
pop()
{
  do{
    ptr = top;            // ptr = top = NodeA
    next_ptr = top->next; // next_ptr = NodeX
  } while(CAS(top, ptr, next_ptr) != true);
  return ptr;   
}
```
而進程P2在執行CAS操作之前打斷了P1，並對棧進行了一系列的pop和push操作，使棧變為如下結構：
```
   top
    |
    V  
  0x0014
| Node C | --> | Node B | --> |  Node X | --> ……
```
進程P2首先pop出NodeA，之後又push了兩個NodeB和C，由於內存管理機制中廣泛使用的內存重用機制，導致NodeC的地址與之前的NodeA一致。

這時P1又開始繼續運行，在執行CAS操作時，由於top依舊指向的是NodeA的地址(實際上已經變為NodeC)，因此將top的值修改為了NodeX，這時棧結構如下：
```
                                   top
                                    |
   0x0014                           V
 | Node C | --> | Node B | --> |  Node X | --> ……
```
經過CAS操作後，top指針錯誤地指向了NodeX而不是NodeB。


### 原子操作

`CAS`的實現需要硬體的支援，通常是一條指令。

如果`CAS`不是原子性,就會相`Test_Lock`一樣出現多個`goroutine`同時進入`critical section`的情況。

[code](./cas_test.go)


```
--- FAIL: Test_Lock (0.00s)
    /go-study/os/process_synchronization/compare_and_swap/cas_test.go:77: num = 978; want 1000
FAIL
FAIL	/go-study/os/process_synchronization/compare_and_swap	0.002s
FAIL
```


### 效能 CAS vs Mutex

[code](./cas_benchmark_test.go)

`CAS`的效能比`Mutex`好一點。

原因是`Mutex`在`Lock`和`Unlock`時都會進行`syscall`，而`CAS`只有在`Lock`時進行`syscall`。

我在讀`golang`的`mutex`的source code時，發現`lock`需要`CompareAndSwapInt32`,而`unlock`需要`AddInt32`。這樣就可以或許解釋為什麼`Mutex`會比`CAS`慢。
> 但`slowlock`我還沒看完，所以這只是我的猜測。


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10875H CPU @ 2.30GHz
BenchmarkIncWithLock
BenchmarkIncWithLock-16          3420673               350.3 ns/op            24 B/op          1 allocs/op
BenchmarkIncWithCAS
BenchmarkIncWithCAS-16           3935949               303.6 ns/op            24 B/op          1 allocs/op
```


### 參考
[並行程式設計: Atomics 操作](https://hackmd.io/@sysprog/concurrency-atomics)
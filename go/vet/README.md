# Vet
> Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.
https://pkg.go.dev/cmd/vet

```shell
go vet my/project/...
```

## 範例

```go
package main

import "fmt"

type nocopy struct{}

func (*nocopy) Lock()   {}
func (*nocopy) Unlock() {}

type TestVet struct {
	nocopy
}

func (t *TestVet) Hello() {
	fmt.Println("Hello")
}

func main() {
	m := TestVet{}
	c := m
	c.Hello()
}

```

這段code直接執行是沒有問題的，但是使用`go vet`檢查後會發現以下問題

```shell
main.go:20:7: assignment copies lock value to c: command-line-arguments.TestVet
```

vet其中一項檢查是:
- copylocks : check for locks erroneously passed by value
 
當我們不想讓struct被copy時,可以嵌套nocopy struct, 這樣就可以被vet檢查出來


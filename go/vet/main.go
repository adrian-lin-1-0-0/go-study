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

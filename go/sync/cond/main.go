package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	fmt.Println(name)

	c.L.Lock()
	for !done {
		fmt.Println(name, "wait")
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()
	fmt.Println("write unlock")
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(new(sync.Mutex))

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)

	fmt.Println("main unlock")
}

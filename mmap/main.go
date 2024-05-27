package main

import (
	"fmt"
	"log"

	"golang.org/x/exp/mmap"
)

func main() {
	at, err := mmap.Open("/tmp/mmap-test")

	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 5)
	_, _ = at.ReadAt(buff, 0)
	_ = at.Close()
	fmt.Println(string(buff))
}

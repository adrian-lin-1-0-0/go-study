package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	sockFileName := "/tmp/test-ipc.sock"

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(proto, addr string) (conn net.Conn, err error) {
				return net.Dial("unix", sockFileName)
			},
		},
	}

	resp, err := client.Get("http://xxx/")
	if err != nil {
		log.Fatal("GET error: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read error: ", err)
	}
	fmt.Println(string(body))
}

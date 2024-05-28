package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sockFileName := "/tmp/test-ipc.sock"

	l, err := net.Listen("unix", sockFileName)
	if err != nil {
		log.Fatal("Listen error: ", err)
	}
	defer l.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(sockFileName)
		os.Exit(1)
	}()

	http.Serve(l,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Received: %s", r.URL.Path)
			fmt.Fprintf(w, "Hello from server")
		}),
	)

}

package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sockFileName := "/tmp/test-ipc.sock"
	socket, err := net.Listen("unix", sockFileName)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(sockFileName)
		os.Exit(1)
	}()

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received: %s", buf[:n])

	_, err = conn.Write([]byte("Hello from server"))
	if err != nil {
		log.Fatal(err)
	}
}

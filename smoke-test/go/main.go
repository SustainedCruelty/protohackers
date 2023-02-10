package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to start the server: %s", err)
	}
	log.Println("listening...")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("failed to accept the connection: %s", err)
		}
		go sendEcho(conn)
	}
}

func sendEcho(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		buffer := make([]byte, 1024)
		n, err := reader.Read(buffer)
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatalf("failed to read: %s", err)
		}
		conn.Write(buffer[:n])
	}
}

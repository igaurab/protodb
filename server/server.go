package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	println("Starting server at :8080")
	var conn_numbers int = 0
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		println("Accepting connection: ", conn_numbers)
		go handleConnection(conn)
		conn_numbers++
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	client := newClient(conn)
	client.read()
}

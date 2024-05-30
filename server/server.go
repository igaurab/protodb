package main

import (
	"bytes"
	"fmt"
	"io"
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

	buf := make([]byte, 1024)
	var requestBuffer bytes.Buffer

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				// End of data
				break
			}
			fmt.Println(err)
			break
		}

		requestBuffer.Write(buf[:n])
	}

	requestStr := requestBuffer.String()

	fmt.Println("Recieved", requestStr)
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	SERVER_URL := "localhost"
	SERVER_PORT := "8080"
	SERVER_ADDR := SERVER_URL + ":" + SERVER_PORT
	conn, err := net.Dial("tcp", SERVER_ADDR)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	fmt.Println("Connected to server at ", SERVER_ADDR)

	for {
		fmt.Print("# ")
		reader := bufio.NewReader(os.Stdin)
		str, _, _ := reader.ReadLine()
		fmt.Println("Client command: ", string(str))
		_, err = conn.Write(str)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func requestServer(conn net.Conn) {

}

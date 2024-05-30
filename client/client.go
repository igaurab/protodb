package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var str = "this is the good times that has been one of the most intresting times of the year so that i have to goto the gym at the very least of the time and i don't want to do that at the meoment"
	_, err = conn.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}
}

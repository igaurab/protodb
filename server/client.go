package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

type client struct {
	conn       net.Conn
	connect    chan<- *client
	register   chan<- *client
	deregister chan<- *client
	db_name    string
}

func newClient(conn net.Conn) *client {
	return &client{
		conn: conn,
	}
}

func (c *client) read() error {

	buf := make([]byte, 1025)
	for {
		_, err := c.conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			fmt.Println(err)
			return err
		}
		c.handle(buf)
	}
}

func (c *client) handle(message []byte) error {

	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))

	switch string(cmd) {

	case "CONN":
		if err := c.reg(args); err != nil {
			c.err(err)
		}
	}

	fmt.Printf("CMD: %s\n", string(cmd))
	fmt.Printf("Args: %s\n", string(args))

	return nil
}

func (c *client) reg(args []byte) error {
	// Register a client with the server
	db_name := bytes.TrimSpace(args)
	db_path := "files" + string(db_name) + ".json"
	err := createDB(db_path)

	if err != nil {
		return err
	}

	c.db_name = db_path
	c.register <- c

	return nil
}

func (c *client) err(e error) {
	c.conn.Write([]byte("ERR " + e.Error() + "\n"))
}

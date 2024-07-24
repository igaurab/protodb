package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

var (
	DELIMITER = []byte(`\r\n`)
)

type client struct {
	conn       net.Conn
	outbound   chan<- command
	register   chan<- *client
	deregister chan<- *client
	db_name    string
}

func newClient(conn net.Conn, o chan<- command, r chan<- *client, d chan<- *client) *client {
	return &client{
		conn:       conn,
		outbound:   o,
		register:   r,
		deregister: d,
	}
}

func (c *client) read() error {

	buf := make([]byte, 1025)
	for {
		_, err := c.conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				c.deregister <- c
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

	fmt.Printf("CMD: %s\n", string(cmd))
	fmt.Printf("Args: %s\n", string(args))

	switch string(cmd) {

	case "CONN":
		if err := c.reg(args); err != nil {
			c.err(err)
		}
	case "CREATE":
		if err := c.create(args); err != nil {
			c.err(err)
		}
	case "READ":
		if err := c.readData(args); err != nil {
			c.err(err)
		}
	case "UPDATE":
		if err := c.update(args); err != nil {
			c.err(err)
		}
	case "DELETE":
		if err := c.delete(args); err != nil {
			c.err(err)
		}
	case "CLOSECONN":
		if err := c.closeConn(args); err != nil {
			c.err(err)
		}
	default:
		c.err(fmt.Errorf("unknown command %s", cmd))
	}

	return nil
}

func (c *client) reg(args []byte) error {
	// Register a client with the server
	db_name := bytes.TrimSpace(args)
	db_path := "files" + string(db_name) + ".json"
	err := createDB(db_path)
	fmt.Println("Down: ", err)

	if err != nil {
		return err
	}

	c.db_name = db_path
	c.register <- c

	fmt.Println("Successfully created connection with db: ", c.db_name)
	return nil
}

func (c *client) create(args []byte) error {
	return nil
}

func (c *client) readData(args []byte) error {
	return nil
}

func (c *client) update(args []byte) error {
	return nil
}

func (c *client) delete(args []byte) error {
	return nil
}

func (c *client) closeConn(args []byte) error {
	return nil
}

func (c *client) err(e error) {
	c.conn.Write([]byte("ERR " + e.Error() + "\n"))
}

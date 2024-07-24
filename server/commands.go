package main

type command struct {
	id        ID
	recipient string
	sender    string
	body      []byte
}

type ID int

const (
	CONN ID = iota
	CREATE
	READ
	UPDATE
	DELETE
	CLOSECONN
)

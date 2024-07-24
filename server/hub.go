package main

import "fmt"

type hub struct {
	clients         map[string]*client
	commands        chan command
	deregistrations chan *client
	registrations   chan *client
}

func newHub() *hub {
	return &hub{
		registrations:   make(chan *client),
		deregistrations: make(chan *client),
		clients:         make(map[string]*client),
		commands:        make(chan command),
	}
}

func (h *hub) run() {
	for {
		select {
		case client := <-h.registrations:
			h.register(client)
		case client := <-h.deregistrations:
			h.deregister(client)
		case cmd := <-h.commands:
			switch cmd.id {
			case CREATE:
				h.createDb()
			case UPDATE:
				h.update()
			case DELETE:
				h.delete()
			case CLOSECONN:
				h.closeConnection()
			case CONN:
				h.connect()
			default:
				// TODO: Handle Error
			}
		}
	}
}

func (h *hub) register(c *client) {
	fmt.Println("Registered client", c.db_name)
	h.clients[c.db_name] = c
}

func (h *hub) deregister(c *client) {
	fmt.Println("deregistered client", c.db_name)

}

func (h *hub) createDb() {
	fmt.Println("deregistered client", h.commands)

}

func (h *hub) update() {
	fmt.Println("deregistered client", h.commands)

}

func (h *hub) delete() {
	fmt.Println("deregistered client", h.commands)

}

func (h *hub) connect() {
	fmt.Println("deregistered client", h.commands)

}

func (h *hub) closeConnection() {
	fmt.Println("deregistered client", h.commands)

}

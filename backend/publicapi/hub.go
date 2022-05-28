package publicapi

import (
	"log"
)

// Hub is a central authority for relaying messages to clients.
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	log        *log.Logger
}

// NewHub creates a new instance of a Hub.
func NewHub(logger *log.Logger) *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		log:        logger,
	}
}

// Run initiates the hub and begins listening for requests.
func (h *Hub) Run() {
	for {
		select {

		// A client has registered. We should include it in our list of clients so that it will recieve
		// messages.
		case client := <-h.register:
			h.log.Printf("registering client: '%s'", client.connection.RemoteAddr().String())
			h.clients[client] = true

		// A client has unregistered. We should remove it from our list of clients so that we are no
		// longer trying to send it any messages.
		case client := <-h.unregister:
			h.log.Printf("unregistering client: '%s'", client.connection.RemoteAddr().String())
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		// A message was sent to the hub. We should broadcast it to all of our clients.
		case message := <-h.broadcast:
			h.log.Printf("broadcasting '%s'", message)

			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

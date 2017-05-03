package main

import (
	"github.com/gorilla/websocket"
	"gochat/trace"
	"log"
	"net/http"
)

const (
	SocketBufferSize  = 1024
	messageBufferSize = 256
)

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	tracer  trace.Tracer
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			//TODO use this if tracing required
			//r.tracer.Trace("Client joined room.")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left room.")
		case msg := <-r.forward:
			r.tracer.Trace("Message received : ", string(msg))
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: SocketBufferSize, WriteBufferSize: SocketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
	}
}

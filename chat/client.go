package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

func (c *client) read() {
	defer c.socket.Close()
	_, msg, err := c.socket.ReadMessage()
	if err != nil {
		log.Fatal(err)
		return
	}
	c.room.forward <- msg
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"encoding/json"
	// "fmt"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// The hub.
	h *hub
}

func (c *connection) reader() {
	for {
		_, event, _ := c.ws.ReadMessage()
		var j miEvento
		// 		c.ws.ReadJSON(j)
		// 		// ReadJSON utiliza encode/json y utiliza las mismas reglas de conversión
		json.Unmarshal(event, &j)

		if j.Action == "broadcast" {
			c.h.broadcast <- []byte(j.Message)
		}
	}
	c.ws.Close()
}

func (c *connection) writer() {
	/* eventualmente tenemos que mandar los mensajes también codificados?
	o no es necesario porque los resultados que se mandan de servidor
	a cliente sólo son cadenas? bueno para el juego no no?*/
	for message := range c.send {
		// por este range de aquí es importante cerrar el canal en hub
		// func (c *Conn) WriteJSON(v interface{}) error
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type wsHandler struct {
	h *hub
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws, h: wsh.h}
	c.h.register <- c
	defer func() { c.h.unregister <- c }()
	go c.writer()
	c.reader()
}

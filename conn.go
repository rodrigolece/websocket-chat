package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan wsEvent

	// Canal no buffereado para que bloquee
	status chan string

	// The hub.
	h *hub

	// El gas
	g *gas
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type wsHandler struct {
	h *hub
	g *gas
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{
		send: make(chan wsEvent, 256),
		status: make(chan string),
		ws: ws,
		h: wsh.h,
		// Cada conexión se asigna a un gas
		g: wsh.g,
	}
	c.h.register <- c
	defer func() { c.h.unregister <- c }()
	go c.writer()
	c.reader()
}

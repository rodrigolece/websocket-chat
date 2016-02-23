package main

type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan wsEvent

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection

	// Las id's de las conexiones
	id2conn map[string]*connection
	conn2id map[*connection]string
}

func newHub() *hub {
	return &hub{
		// 10: número arbitrario, pero debe servir como colchón para que
		// broadcast no se llene y se bloquee
		broadcast:   make(chan wsEvent),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),
		id2conn:	 make(map[string]*connection),
		conn2id:	 make(map[*connection]string),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			go register(h, c)
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
				// close es importante porque termina la lectura en loops de range
				id := h.conn2id[c]
				delete(h.id2conn, id)
				delete(h.conn2id, c)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m: // falla si el buffer está lleno
				default:
					h.unregister <- c
				}
			}
		}
	}
}

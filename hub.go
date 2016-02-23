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
		// número arbitrario, pero debe servir como colchón
		broadcast:   make(chan wsEvent, 10),
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
			id := randString(5)
			// Anunciamos al nuevo miembro...
			newMember := wsEvent{
				Action: "register",
				Data: id,
			}
			// Para canales buffereados, no podemos asegurar el orden de eventos:
			// el registro termina y LUEGO se lee broadcast.
			c.h.broadcast <- newMember
			// select {
			// case c.h.broadcast <- newMember:
			// default:
			// }
			//... y luego lo registramos
			h.connections[c] = true
			h.id2conn[id] = c
			h.conn2id[c] = id
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

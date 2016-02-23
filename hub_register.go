package main

func register(h *hub, c *connection) {
    ownID := randString(5)
    c.send <- wsEvent{
        Action: "registerown",
        Data: ownID,
    }
    // Necesitamos esperar la respuesta para sincronizar los eventos
    if "ok" == <- c.status {
        // El resto de los clientes registra al nuevo cliente
        c.h.broadcast <- wsEvent{
            Action: "register",
            Data: ownID,
        }
        // El nuevo cliente registra a todos los clientes ya conectados
        for id := range h.id2conn {
            c.send <- wsEvent{
                Action: "register",
                Data: id,
            }
        }
        // Finalmente actualizamos los maps para incluir al nuevo cliente
        h.connections[c] = true
        h.id2conn[ownID] = c
        h.conn2id[c] = ownID
    }
}

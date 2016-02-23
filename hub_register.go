package main

// import (
//     "fmt"
// )

func register(h *hub, c *connection) {
    id := randString(5)
    c.send <- wsEvent{
        Action: "registerown",
        Data: id,
    }
    if "ok" == <- c.status {
        c.h.broadcast <- wsEvent{
            Action: "register",
            Data: id,
        }
        h.connections[c] = true
        h.id2conn[id] = c
        h.conn2id[c] = id
    }
    // Para canales buffereados, no podemos asegurar el orden de eventos:
    // el registro termina y LUEGO se lee broadcast.
    // select {
    // case c.h.broadcast <- newMember:
    // default:
    // }

}

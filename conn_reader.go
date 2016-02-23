package main

import (
    "encoding/json"
)

func (c *connection) reader() {
	for {
		_, event, _ := c.ws.ReadMessage()
		var j wsEvent
		// 		c.ws.ReadJSON(j)
		// 		// ReadJSON utiliza encode/json y utiliza las mismas reglas de conversión
		json.Unmarshal(event, &j)
		go handleWsEvent(c, j)
	}
	c.ws.Close()
}

func handleWsEvent(c *connection, j wsEvent) {
	id := c.h.conn2id[c]

    event := wsEvent{
        Action: "",
        Message: id + ": "+ j.Message,
    }

	switch j.Action {
	case "broadcast":
		c.h.broadcast <- event
	case "sendto":
		recipient, ok := j.Data.(string) // La id del destinatario
		if !ok { return }
		if recipientConn, ok := c.h.id2conn[recipient]; ok {
			recipientConn.send <- event
			c.send <- event
		}
    case "registerownOK":
        c.status <- "ok"
	}
}

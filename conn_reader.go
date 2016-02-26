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

    var event wsEvent
    Data, ok := j.Data.(data)
    var DataArray []data
    if !ok {
        DataArray, _ = j.Data.([]data)
    }

	switch j.Action {
	case "broadcast":
        if Data.Type == "message" {
            event = wsEvent{
                Action: "read",
                Data: data{
                    Type: "message",
                    Content: id + ": "+ Data.Content.(string),
                },
            }
            c.h.broadcast <- event
        }
        if Data.Type == "turn" {
            // direction := j.Data.Content
        }
        if Data.Type == "stopturn" {

        }
	case "sendto":
        // idMessage, ok := j.Data.([]string)
		recipient, ok := DataArray[0].Content.(string) // La id del destinatario
		if !ok { return }
		if recipientConn, ok := c.h.id2conn[recipient]; ok {
            event = wsEvent{
                Action: "read",
                Data: data{
                    Type: "message",
                    Content: id + ": "+ DataArray[1].Content.(string),
                },
            }
			recipientConn.send <- event
			c.send <- event
		}
    case "registerownOK":
        // Para sincronizar el registro de un nuevo cliente
        c.status <- "ok"
	}
}

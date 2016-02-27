package main

import (
    "fmt"
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
    eventData := make([]data, 0)
    // Data, ok := j.Data.([]data)

    d := j.Data[0] // el primer elemento; la mayoría de la veces el único

	switch j.Action {
	case "broadcast":
        if d.Type == "message" {
            eventData = append(eventData, data{
                Type: "message",
                Content: id + ": "+ d.Content.(string),
            })
            event = wsEvent{
                Action: "readmessage",
                Data: eventData,
            }
            c.h.broadcast <- event
        }
        // Si recibimos la posisión de una partícula se la mandamos al gas
        if d.Type == "pos"{
            e := j.Data[1] // para pos hay dos elementos en Data
            pos := d.Content.(map[string]interface{})
            x := pos.x.(float64)
            fmt.Println(x)
            // part := &particle{
            //     pos: d.Content.(map[string]float64),
            //     vel: e.Content.(map[string]float64),
            // }
            // c.g.register <- part
        }
        if d.Type == "turn" {
            // direction := d.Content
        }
        if d.Type == "stopturn" {

        }
	case "sendto":
        e := j.Data[1] // para sendto el segundo elemento es el mensaje
        // el primero la id
		recipient, ok := d.Content.(string) // La id del destinatario
		if !ok { return }
		if recipientConn, ok := c.h.id2conn[recipient]; ok {
            eventData = append(eventData, data{
                Type: "message",
                Content: id + ": "+ e.Content.(string),
            })
            event = wsEvent{
                Action: "readmessage",
                Data: eventData,
            }
			recipientConn.send <- event
			c.send <- event
		}
    case "registerownOK":
        // Para sincronizar el registro de un nuevo cliente
        c.status <- "ok"
	}
}

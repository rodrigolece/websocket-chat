package main

import (
    "github.com/gorilla/websocket"
    "encoding/json"
)

func (c *connection) writer() {

	for event := range c.send {
		// por este range de aquí es importante cerrar el canal en hub
        // Convertimos el evento a una cadena de json
        j, _ := json.Marshal(event)
        // Y lo mandamos después de convertilo a array de bytes
		err := c.ws.WriteMessage(websocket.TextMessage, []byte(j))
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

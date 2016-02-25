package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

var (
	addr      = ":8080"
	client    = "./home.html"
)

func homeHandler(c http.ResponseWriter, req *http.Request) {
	toByte, _ := ioutil.ReadFile(client)
	c.Write(toByte)
}

func main() {
	h := newHub()
	go h.run()
	// Registra una función handler
	http.HandleFunc("/", homeHandler)
	// Registra un Handler (interfaz que implementa el método ServeHTTP)
	http.Handle("/ws", wsHandler{h: h})
	// El script para dibujar canvas
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static", http.StripPrefix("/static", fs))
	// ----
	// http.Handle("/static", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/static/gas.js", func(c http.ResponseWriter, req *http.Request) {
        http.ServeFile(c, req, "./static/gas.js")
	})
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

package main

import (
	"flag"
	"log"
	"net/http"
)

type Server struct {
	// Registered hub.
	hubs map[string]*Hub
}

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func findOrCreateHub(server *Server, addr string) *Hub {
	if hub, ok := server.hubs[addr]; ok {
		return hub
	}

	hub := newHub()
	server.hubs[addr] = hub
	return hub
}

func main() {
	flag.Parse()
	server := &Server{hubs: make(map[string]*Hub)}

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wantedRoom := r.URL.Query().Get("room")

		hub := findOrCreateHub(server, wantedRoom)
		go hub.run()

		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

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
var server = &Server{hubs: make(map[string]*Hub)}

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

func findOrCreateHub(server *Server, r *http.Request) *Hub {
	wantedRoom := r.URL.Query().Get("room")

	if hub, ok := server.hubs[wantedRoom]; ok {
		print("found hub")
		return hub
	}
	print("create")
	hub := newHub()
	server.hubs[wantedRoom] = hub
	return hub
}

func main() {
	flag.Parse()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", addUserToRoom)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func addUserToRoom(w http.ResponseWriter, r *http.Request) {
	hub := findOrCreateHub(server, r)

	go hub.run()

	serveWs(hub, w, r)
}

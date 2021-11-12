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
	http.HandleFunc("/ws", start)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	oldLenght := len(server.hubs)
	hub := findOrCreateHub(server, r)
	newLenght := len(server.hubs)

	if newLenght > oldLenght {
		// I really don't get it, if the hub is created we can't access the broadcast channel
		// the only solution I currently find is to re execute the function
		// look like the hub must be create before calling the start function
		start(w, r)
	}

	go hub.run()

	serveWs(hub, w, r)
}

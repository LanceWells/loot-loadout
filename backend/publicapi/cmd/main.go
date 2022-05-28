package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/lantspants/lootloadout/publicapi"
)

// addr is the service address that this socket will listen on.
var addr = flag.String("addr", ":8080", "http service address")

// serveHome provides a webpage that works with this chat socket application (this will be removed).
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./service_public/home.html")
}

// main is the main function for this application.
func main() {
	log := log.New(os.Stdout, "", 0)

	flag.Parse()
	hub := publicapi.NewHub(log)
	go hub.Run()

	log.Printf("Initializing WS service at %v", *addr)

	http.HandleFunc("/", serveHome)

	// Once a client hits this endpoint, serve it a websocket client so that it may listen to the hub
	// that is broadcasting messages.
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		publicapi.ServeWS(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

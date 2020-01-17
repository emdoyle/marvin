package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Started")
	staticFileServer := http.FileServer(http.Dir(GetEnv("STATIC_DIR", "assets/build/")))

	http.Handle("/", staticFileServer)
	http.HandleFunc("/events", EventHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

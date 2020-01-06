package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Started")

	staticFileServer := http.FileServer(http.Dir("assets/build/"))

	http.Handle("/", staticFileServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

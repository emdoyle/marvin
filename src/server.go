package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Print("Started")
	staticFileServer := http.FileServer(http.Dir(GetEnv("STATIC_DIR", "assets/build/")))

	router := mux.NewRouter()
	router.Handle("/", staticFileServer)
	router.HandleFunc("/events", EventHandler)
	router.HandleFunc("/challenge", ChallengeHandler)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

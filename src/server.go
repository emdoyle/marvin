package main

import (
	"github.com/emdoyle/marvin/src/models"
	"log"
	"net/http"
)

func main() {
	log.Print("Started")
	log.Print("Running migrations...")
	DB.AutoMigrate(&models.User{})
	log.Print("Migrations ran.")
	staticFileServer := http.FileServer(http.Dir(GetEnv("STATIC_DIR", "assets/build/")))

	http.Handle("/", staticFileServer)
	http.HandleFunc("/events", EventHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

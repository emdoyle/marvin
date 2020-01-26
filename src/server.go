package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emdoyle/marvin/src/models"
)

func migrate() {
	if DB != nil {
		log.Print("Running migrations...")
		DB.AutoMigrate(&models.User{})
		log.Print("Migrations ran.")
	} else {
		log.Print("Not running migrations since DB unavailable.")
	}
}

func main() {
	log.Print("Started")
	migrate()

	staticFileServer := http.FileServer(http.Dir(GetEnv("STATIC_DIR", "assets/build/")))
	http.Handle("/", staticFileServer)
	http.HandleFunc("/events", EventHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", GetEnv("MARVIN_PORT", "8080")), nil))
}

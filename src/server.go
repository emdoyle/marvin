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
	http.HandleFunc("/interactive", InteractionHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", GetEnv("MARVIN_PORT", "8080")), nil))
}

//SetUpJSONResponse configures a ResponseWriter to send JSON
func SetUpJSONResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
}

//DeclineResponse sends a 403 response using a ResponseWriter
func DeclineResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusUnauthorized)
	writer.Write([]byte("denied"))
}

//VerifySlackSignature takes an http Request and verifies that it was
//signed by Slack
func VerifySlackSignature(rawBody []byte, request *http.Request) bool {
	timestamp := request.Header.Get(SlackTimestampHeader)
	slackSignature := []byte(request.Header.Get(SlackSignatureHeader))
	return VerifySigningSignature(timestamp, rawBody, slackSignature)
}

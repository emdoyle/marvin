package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	http.Handle("/events", WithSlackSignatureVerificationHandler(EventHandler))
	http.Handle("/interactive", WithSlackSignatureVerificationHandler(InteractionHandler))
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
	log.Printf("Denied request")
}

//FailResponse sends a 500 response using a ResponseWriter
func FailResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Write([]byte("error"))
	log.Printf("Failed request")
}

//VerifySlackSignature takes an http Request and verifies that it was
//signed by Slack
func VerifySlackSignature(rawBody []byte, request *http.Request) bool {
	timestamp := request.Header.Get(SlackTimestampHeader)
	slackSignature := []byte(request.Header.Get(SlackSignatureHeader))
	return VerifySigningSignature(timestamp, rawBody, slackSignature)
}

//GetRawBody reads the request body and replaces the Body with a new Reader
func GetRawBody(request *http.Request) ([]byte, error) {
	rawBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return []byte{}, err
	}
	request.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
	return rawBody, nil
}

//WithSlackSignatureVerificationHandler decorates a HandlerFunc with logic to
//verify the request came from Slack
func WithSlackSignatureVerificationHandler(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawBody, err := GetRawBody(r)
		if err != nil {
			log.Printf("Failed to get body from incoming request")
			FailResponse(w)
			return
		}
		if verified := VerifySlackSignature(rawBody, r); !verified {
			DeclineResponse(w)
			return
		}
		handler(w, r)
	})
}

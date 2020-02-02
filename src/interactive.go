package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

//InteractionHandler handles interaction payloads from Slack
func InteractionHandler(w http.ResponseWriter, r *http.Request) {
	rawBody, _ := ioutil.ReadAll(r.Body)
	if verified := VerifySlackSignature(rawBody, r); !verified {
		DeclineResponse(w)
		return
	}
	log.Printf("Handle interaction")
}

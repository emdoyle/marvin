package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//UserObject is a JSON payload with user info
type UserObject struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	TeamID   string `json:"team_id"`
}

//InteractionPayload is a JSON payload describing a user interaction
type InteractionPayload struct {
	Type        string      `json:"type"`
	TriggerID   string      `json:"trigger_id"`
	ResponseURL string      `json:"response_url"`
	User        UserObject  `json:"user"`
	MessageSrc  interface{} `json:"message,omitempty"`
	ViewSrc     interface{} `json:"view,omitempty"`
	Actions     interface{} `json:"actions"`
}

func handleBlockActions(payload InteractionPayload) {
	log.Printf("Handling block_actions payload")
}

//InteractionHandler handles interaction payloads from Slack
func InteractionHandler(w http.ResponseWriter, r *http.Request) {
	rawBody, err := GetRawBody(r)
	if err != nil {
		log.Printf("Failed to get body from interaction request")
		return
	}
	if verified := VerifySlackSignature(rawBody, r); !verified {
		DeclineResponse(w)
		return
	}

	r.ParseForm()
	payload, present := r.Form["payload"]
	if !present {
		log.Printf("Could not find payload in interaction form data!")
		return
	}
	var interactionPayload InteractionPayload
	err = json.Unmarshal(([]byte)(payload[0]), &interactionPayload)
	if err != nil {
		log.Printf("Error in unmarshal: %s", err)
		return
	}

	log.Printf("Received interaction payload: %v", interactionPayload)
	switch interactionPayload.Type {
	case "block_actions":
		handleBlockActions(interactionPayload)
	default:
		log.Printf("Received an unknown interaction payload")
	}
}

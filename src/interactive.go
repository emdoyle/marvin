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

//InteractionResponse is a JSON response sent to the ResponseURL of an Interaction
type InteractionResponse struct {
	Text            string `json:"text,omitempty"`
	ResponseType    string `json:"response_type,omitempty"`
	ReplaceOriginal bool   `json:"replace_original,omitempty"`
	DeleteOriginal  bool   `json:"delete_original,omitempty"`
}

//ActionObject is a JSON object describing a single action on an app surface
type ActionObject struct {
	BlockID  string     `json:"block_id"`
	ActionID string     `json:"action_id"`
	ActionTS string     `json:"action_ts,omitempty"`
	Text     TextObject `json:"text"`
}

//InteractionPayload is a JSON payload describing a user interaction
type InteractionPayload struct {
	Type        string         `json:"type"`
	TriggerID   string         `json:"trigger_id"`
	ResponseURL string         `json:"response_url"`
	User        UserObject     `json:"user"`
	MessageSrc  interface{}    `json:"message,omitempty"`
	ViewSrc     interface{}    `json:"view,omitempty"`
	Actions     []ActionObject `json:"actions"`
}

func handleBlockActions(payload InteractionPayload) {
	log.Printf("Handling block_actions payload")
	for _, action := range payload.Actions {
		currResponse := InteractionResponse{
			Text:            action.Text.Text,
			ReplaceOriginal: true,
		}
		POSTToURL(currResponse, payload.ResponseURL)
	}
}

//InteractionHandler handles interaction payloads from Slack
func InteractionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload, present := r.Form["payload"]
	if !present || len(payload) == 0 {
		log.Printf("Could not find payload in interaction form data!")
		FailResponse(w)
		return
	}
	var interactionPayload InteractionPayload
	err = json.Unmarshal(([]byte)(payload[0]), &interactionPayload)
	if err != nil {
		log.Printf("Error in unmarshal: %s", err)
		FailResponse(w)
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

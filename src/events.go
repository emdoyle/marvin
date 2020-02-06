package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//EventWrapper is the top-level JSON payload
type EventWrapper struct {
	Token       string   `json:"token"`
	TeamID      string   `json:"team_id"`
	APIAppID    string   `json:"api_app_id"`
	Event       *Event   `json:"event"`
	AuthedUsers []string `json:"authed_users"`
	EventID     string   `json:"event_id"`
	EventTime   int64    `json:"event_time"`
	Challenge   string   `json:"challenge"`
	Type        string   `json:"type"`
}

//Event is the JSON structure of a single event
type Event struct {
	Type        string `json:"type"`
	Timestamp   string `json:"event_ts"`
	User        string `json:"user"`
	Channel     string `json:"channel"`
	ChannelType string `json:"channel_type"`
	Text        string `json:"text"`
}

//ChallengeResponse is the response we send back for a challenge event
type ChallengeResponse struct {
	Challenge string `json:"challenge"`
}

func handleChallenge(eventWrapper EventWrapper, writer http.ResponseWriter) {
	response := ChallengeResponse{
		Challenge: eventWrapper.Challenge,
	}
	SetUpJSONResponse(writer)
	json.NewEncoder(writer).Encode(response)
}

func handleEvent(event Event) {
	switch event.Type {
	case "message":
		HandleMessage(event)
	case "app_mention":
		HandleMention(event)
	}
}

//EventHandler handles Slack's events
func EventHandler(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		FailResponse(w)
		return
	}
	var eventWrapper EventWrapper
	json.Unmarshal(rawBody, &eventWrapper)

	switch eventWrapper.Type {
	case "url_verification":
		handleChallenge(eventWrapper, w)
	case "event_callback":
		handleEvent(*eventWrapper.Event)
	}
}

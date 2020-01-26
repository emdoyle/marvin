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
	Event       Event    `json:"event"`
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

func setUpJSONResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func declineResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusUnauthorized)
	writer.Write([]byte("denied"))
}

func handleChallenge(eventWrapper EventWrapper, writer http.ResponseWriter) {
	response := ChallengeResponse{
		Challenge: eventWrapper.Challenge,
	}
	setUpJSONResponse(writer)
	json.NewEncoder(writer).Encode(response)
}

func verifySlackSignature(rawBody []byte, request *http.Request) bool {
	timestamp := request.Header.Get("X-Slack-Request-Timestamp")
	slackSignature := []byte(request.Header.Get("X-Slack-Signature"))
	return VerifySigningSignature(timestamp, rawBody, slackSignature)
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
	rawBody, _ := ioutil.ReadAll(r.Body)
	if verified := verifySlackSignature(rawBody, r); !verified {
		declineResponse(w)
		return
	}
	var eventWrapper EventWrapper
	json.Unmarshal(rawBody, &eventWrapper)

	switch eventWrapper.Type {
	case "url_verification":
		handleChallenge(eventWrapper, w)
	case "event_callback":
		handleEvent(eventWrapper.Event)
	}
}

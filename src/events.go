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
	Type      string `json:"type"`
	Timestamp string `json:"event_ts"`
	User      string `json:"user"`
}

//ChallengeResponse is the response we send back for a challenge event
type ChallengeResponse struct {
	Challenge string `json:"challenge"`
}

func setUpJSONResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
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

func verifySlackSignature(request *http.Request) bool {
	rawBody, _ := ioutil.ReadAll(request.Body)
	timestamp := request.Header.Get("X-Slack-Request-Timestamp")
	slackSignature := []byte(request.Header.Get("X-Slack-Signature"))
	return VerifySigningSignature(timestamp, rawBody, slackSignature)
}

//EventHandler handles Slack's events
func EventHandler(w http.ResponseWriter, r *http.Request) {
	if verified := verifySlackSignature(r); !verified {
		declineResponse(w)
		return
	}
	var eventWrapper EventWrapper
	json.NewDecoder(r.Body).Decode(&eventWrapper)

	if eventWrapper.Type == "url_verification" {
		handleChallenge(eventWrapper, w)
	}
}

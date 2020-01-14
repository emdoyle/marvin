package main

import (
	"encoding/json"
	"net/http"
)

//Challenge is the payload sent by Slack
type Challenge struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

//Response is the response we send back
type Response struct {
	Challenge string `json:"challenge"`
}

//ChallengeHandler handles Slack's verification request
func ChallengeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: verify Slack request signature
	var challenge Challenge
	json.NewDecoder(r.Body).Decode(&challenge)

	response := Response{
		Challenge: challenge.Challenge,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

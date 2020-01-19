package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Message is a message payload
type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func setJSONResponse(request *http.Request) {
	request.Header.Set("Accept-Charset", "utf-8")
	request.Header.Set("Content-Type", "application/json")
}

func addAuthToken(request *http.Request) {
	request.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", GetEnv("SLACKBOT_OAUTH_TOKEN", "")),
	)
}

//POSTToSlack takes a Message and makes an authorized POST request
func POSTToSlack(message *Message) {
	client := &http.Client{}
	payload, _ := json.Marshal(message)
	log.Printf("%s", payload)
	request, err := http.NewRequest(
		"POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(payload),
	)
	if err != nil {
		log.Printf("Error creating POST request: %s", err)
		return
	}
	addAuthToken(request)
	setJSONResponse(request)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending POST request: %s", err)
	}
	rawBody, _ := ioutil.ReadAll(response.Body)
	log.Printf(
		"Response status: %s statusCode: %v body: %s",
		response.Status,
		response.StatusCode,
		string(rawBody),
	)
}

//HandleMessage handles a 'message' type Event
func HandleMessage(event Event) {
	log.Printf("Received message text: %s", event.Text)
}

//HandleMention handles an 'app_mention' type Event
func HandleMention(event Event) {
	log.Printf("Received mention text: %s", event.Text)
	message := &Message{
		Channel: event.Channel,
		Text:    "Don't bother me.",
	}
	POSTToSlack(message)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Message is a message payload
type Message struct {
	Channel  string      `json:"channel"`
	Text     string      `json:"text"`
	Blocks   interface{} `json:"blocks"`
	ThreadTs string      `json:"thread_ts"`
	Markdown bool        `json:"mrkdwn"`
}

func setJSONResponse(request *http.Request) {
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
}

func addAuthToken(request *http.Request) {
	request.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", GetEnv("SLACKBOT_OAUTH_TOKEN", "")),
	)
}

//POSTToURL takes a JSON-serializable parameter
//and makes an auth'd POST request to a URL
func POSTToURL(message interface{}, url string) error {
	client := &http.Client{}
	payload, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error creating POST request: %s", err)
		return err
	}
	log.Printf("Posting payload: %s", payload)
	request, err := http.NewRequest(
		"POST", url, bytes.NewBuffer(payload),
	)
	if err != nil {
		log.Printf("Error creating POST request: %s", err)
		return err
	}
	addAuthToken(request)
	setJSONResponse(request)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending POST request: %s", err)
		return err
	}
	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading POST response: %s", err)
		return err
	}
	log.Printf(
		"Response status: %s statusCode: %v body: %s",
		response.Status,
		response.StatusCode,
		string(rawBody),
	)
	return nil
}

//POSTToSlack takes a Message and makes an authorized POST request to
//the default Slack URL to post chat messages
func POSTToSlack(message interface{}) error {
	err := POSTToURL(message, SlackChatPostMessageURL)
	return err
}

//HandleMessage handles a 'message' type Event
func HandleMessage(event Event) {
	log.Printf("Received message text: %s", event.Text)
	switch event.ChannelType {
	case "im":
		log.Print("Handling API request in IM")
		HandleUserAPIRequest(event)
	}
}

func postInteractiveStart(event Event) error {
	log.Printf("Posting interactive start")
	options := []string{"Pizza", "Sushi", "Pasta"}

	blocks := make([]interface{}, 2)
	blocks[0] = (interface{})(BuildBasicSection("*Pick an option.*"))
	blocks[1] = (interface{})(BuildBasicActions(options))

	message := Message{
		Channel: event.Channel,
		Text:    "Could not display interactive buttons!",
		Blocks:  (interface{})(blocks),
	}
	return POSTToSlack(message)
}

func postInteractiveStop(event Event) error {
	log.Printf("Posting interactive stop")

	message := Message{
		Channel: event.Channel,
		Text:    "Stopping interactivity",
	}
	return POSTToSlack(message)
}

func postDefaultMessage(event Event) error {
	message := Message{
		Channel: event.Channel,
		//Text is a fallback when Blocks is passed
		Text: "Don't bother me.",
		Blocks: (interface{})([]SectionBlock{
			BuildBasicSection("Here's a _section_ for you..."),
		}),
	}
	return POSTToSlack(message)
}

//HandleMention handles an 'app_mention' type Event
func HandleMention(event Event) {
	log.Printf("Received mention text: %s", event.Text)
	switch {
	case strings.Contains(event.Text, "start"):
		postInteractiveStart(event)
	case strings.Contains(event.Text, "stop"):
		postInteractiveStop(event)
	default:
		postDefaultMessage(event)
	}
}

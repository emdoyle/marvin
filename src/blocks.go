package main

import (
	"fmt"
	"log"
)

//TextObject represents formatted text, and is composed in various types of Blocks
type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

//ObjectFromText generates a full TextObject from an input string
func ObjectFromText(text string) *TextObject {
	return &TextObject{
		Type:     "mrkdwn",
		Text:     text,
		Verbatim: false,
	}
}

//SectionBlock is a Block Kit component with type 'section'
type SectionBlock struct {
	Type    string        `json:"type"`
	Text    *TextObject   `json:"text"`
	BlockID string        `json:"block_id,omitempty"`
	Fields  []*TextObject `json:"fields,omitempty"`
	//Accessory can be a variable shape, and must be inspected before decoding into another Struct
	Accessory interface{} `json:"accessory,omitempty"`
}

//ActionsBlock is a Block Kit component with type 'actions'
type ActionsBlock struct {
	Type     string      `json:"type"`
	Elements interface{} `json:"elements"`
	BlockID  string      `json:"block_id,omitempty"`
}

//Button is a Block Kit interactive component with type 'button'
type Button struct {
	Type     string        `json:"type"`
	Text     *TextObject   `json:"text"`
	ActionID string        `json:"action_id"`
	URL      string        `json:"url,omitempty"`
	Value    string        `json:"value,omitempty"`
	Style    string        `json:"style,omitempty"`
	Confirm  *Confirmation `json:"confirm,omitempty"`
}

//Confirmation is a Block Kit interactive component for a confirmation modal
type Confirmation struct {
	//Title is plain_text only
	Title TextObject `json:"title"`
	Text  TextObject `json:"text"`
	//Confirm is plain_text only
	Confirm TextObject `json:"confirm"`
	//Deny is plain_text only
	Deny TextObject `json:"deny"`
}

//ButtonsFromOptions makes a slice of Buttons with string options
func ButtonsFromOptions(options []string) []Button {
	results := make([]Button, len(options))
	for i, option := range options {
		results[i] = Button{
			Type: "button",
			Text: &TextObject{
				Type: "plain_text",
				Text: option,
			},
			Value: fmt.Sprintf("%v", i),
		}
	}
	return results
}

//BuildBasicSection just jams the text into an unadorned Section
func BuildBasicSection(text string) SectionBlock {
	log.Printf("Building basic section block")
	return SectionBlock{
		Type: "section",
		Text: ObjectFromText(text),
	}
}

//BuildBasicActions creates an ActionsBlock with buttons corresponding to
//string options
func BuildBasicActions(options []string) ActionsBlock {
	log.Printf("Building basic actions block")
	return ActionsBlock{
		Type:     "actions",
		Elements: (interface{})(ButtonsFromOptions(options)),
	}
}

package main

//TextObject represents formatted text, and is composed in various types of Blocks
type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

//ObjectFromText generates a full TextObject from an input string
func ObjectFromText(text string) TextObject {
	return TextObject{
		Type:     "mrkdwn",
		Text:     text,
		Verbatim: false,
	}
}

//SectionBlock is a Block Kit component with type 'section'
type SectionBlock struct {
	Type    string       `json:"type"`
	Text    TextObject   `json:"text"`
	BlockID string       `json:"block_id,omitempty"`
	Fields  []TextObject `json:"fields,omitempty"`
	//Accessory can be a variable shape, and must be inspected before decoding into another Struct
	Accessory interface{} `json:"accessory,omitempty"`
}

//BuildBasicSection just jams the text into an unadorned Section
func BuildBasicSection(text string) SectionBlock {
	return SectionBlock{
		Type: "section",
		Text: ObjectFromText(text),
	}
}

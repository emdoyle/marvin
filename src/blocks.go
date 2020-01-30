package main

//TextObject represents formatted text, and is composed in various types of Blocks
type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji"`
	Verbatim bool   `json:"verbatim"`
}

//SectionBlock is a Block Kit component with type 'section'
type SectionBlock struct {
	Type    string       `json:"type"`
	Text    TextObject   `json:"text"`
	BlockID string       `json:"block_id"`
	Fields  []TextObject `json:"fields"`
	//Accessory can be a variable shape, and must be inspected before decoding into another Struct
	Accessory interface{} `json:"accessory"`
}

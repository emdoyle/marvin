package main

//TextObject represents formatted text, and is composed in various types of Blocks
type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji"`
	Verbatim bool   `json:"verbatim"`
}

//Block is a Block Kit component
type Block struct {
	Type    string       `json:"type"`
	Text    TextObject   `json:"text"`
	BlockID string       `json:"block_id"`
	Fields  []TextObject `json:"fields"`
	//Accessory can be a variable shape, and must be inspected before decoding into another Struct
	Accessory map[string]interface{} `json:"accessory"`
}

package models

type SlackMessage struct {
	Ok       bool      `json:"ok"`
	Latest   string    `json:"latest"`
	Messages []Message `json:"messages"`
	HasOne   bool      `json:"has_one"`
	Error    string    `json:"error"`
}

type Message struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
	Ts   string `json:"ts"`
}

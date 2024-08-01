package models

type SlackInteractiveResponse struct {
	ResponseType    string `json:"response_type"`
	ReplaceOriginal bool   `json:"replace_original"`
	Text            string `json:"text"`
}

type SlackInteractivePayload struct {
	Type            string         `json:"type"`
	CallbackId      string         `json:"callback_id"`
	OriginalMessage MessageContent `json:"original_message"`
	Action          []ActionDetail `json:"actions"`
	User            SlackUser      `json:"user"`
	Channel         SlackChannel   `json:"channel"`
}

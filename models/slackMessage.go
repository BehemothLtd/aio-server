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

type MessageAttachment struct {
	Text           string           `json:"text"`
	CallbackId     string           `json:"callback_id"`
	Color          string           `json:"color"`
	AttachmentType string           `json:"attachment_type"`
	Actions        []AttachedAction `json:"actions"`
}

type AttachedAction struct {
	Name    string       `json:"name"`
	Text    string       `json:"text"`
	Type    string       `json:"type"`
	Value   string       `json:"value"`
	Confirm ActionDetail `json:"confirm"`
}

type ActionDetail struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	OkText      string `json:"ok_text"`
	DismissText string `json:"dismiss_text"`
}

func NewMessageAttachment(callback string) *[]MessageAttachment {
	switch callback {
	case "change_state_rq":
		return ChangeStateRequest()
	default:
		return nil
	}
}

func ChangeStateRequest() *[]MessageAttachment {
	attachments := []MessageAttachment{
		{
			Text:           "Descision",
			CallbackId:     "change_state_rq",
			Color:          "#3AA3E3",
			AttachmentType: "default",
			Actions: []AttachedAction{
				{
					Name:  "approve",
					Text:  "Approve",
					Type:  "button",
					Value: "approved",
					Confirm: ActionDetail{
						Title:       "Are you sure?",
						Text:        "Approve",
						OkText:      "Yes",
						DismissText: "No",
					},
				},
				{
					Name:  "reject",
					Text:  "Reject",
					Type:  "button",
					Value: "rejected",
					Confirm: ActionDetail{
						Title:       "Are you sure?",
						Text:        "Reject",
						OkText:      "Yes",
						DismissText: "No",
					},
				},
			},
		},
	}

	return &attachments
}

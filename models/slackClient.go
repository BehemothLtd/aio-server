package models

import (
	"aio-server/exceptions"
	"aio-server/pkg/constants"
	"aio-server/pkg/utilities"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type SlackClient struct {
	Client   *http.Client
	EndPoint string
	Headers  map[string]string
}

func NewSlackClient() *SlackClient {
	client := SlackClient{}

	httpClient := http.Client{
		Timeout: constants.RequestTimeOut * time.Second,
	}
	client.Client = &httpClient
	client.EndPoint = os.Getenv("SLACK_MESSAGE_API_ENDPOINT")
	client.Headers = map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
	}

	return &client
}

func (client *SlackClient) Request(method string, endpoint string, payload []byte) (*http.Request, error) {
	requestEndPoint := client.EndPoint + endpoint

	if request, err := http.NewRequest(method, requestEndPoint, bytes.NewBuffer(payload)); err != nil {
		return nil, err
	} else {
		// Set request headers
		additionHeaders := map[string]string{
			"Authorization": fmt.Sprintf("Bearer %+v", os.Getenv("SLACK_BOT_TOKEN")),
		}

		client.SetHeaders(&additionHeaders)

		for key, value := range client.Headers {
			request.Header.Set(key, value)
		}

		return request, nil
	}
}

func (client *SlackClient) FetchConversationHistories(channel string, limit *int, timestamp *time.Time, inclusion *bool) (*SlackMessage, error) {
	if limit == nil {
		num := 20
		limit = &num
	}
	if timestamp == nil {
		now := time.Now()
		timestamp = &now
	}
	if inclusion == nil {
		defaultVal := true
		inclusion = &defaultVal
	}

	payload := map[string]string{
		"channel":   channel,
		"limit":     utilities.IntToString(*limit),
		"latest":    utilities.UnixTimestamp(*timestamp),
		"inclusion": utilities.BoolToString(*inclusion),
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	request, err := client.Request(constants.Post, "/conversations.history", payloadBytes)

	if err != nil {
		return nil, err
	}

	response, err := client.Client.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	message := SlackMessage{}
	err = json.Unmarshal(body, &message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (client *SlackClient) SetHeaders(additionHeaders *map[string]string) *SlackClient {
	if additionHeaders != nil {
		for key, value := range *additionHeaders {
			client.Headers[key] = value
		}
	}

	return client
}

func (client *SlackClient) SendMessage(text string, channel string, callback *string, threadTs *string) (*SlackMessage, error) {
	if text == "" || channel == "" {
		return nil, exceptions.NewBadRequestError("Text and channel are required")
	}

	fmt.Print("\n\n============================\nStart send_message")

	payload := map[string]interface{}{
		"text":    text,
		"channel": channel,
	}

	// Set message's attachment
	if callback != nil {
		attachment := NewMessageAttachment(*callback)

		if _, err := json.Marshal(attachment); err != nil {
			return nil, exceptions.NewBadRequestError("invalid attachments")
		} else {
			payload["attachments"] = *attachment
		}
	}

	// In case of reply to a message's thread
	if threadTs != nil {
		payload["thread_ts"] = *threadTs
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	endPoint := "/chat.postMessage"
	request, err := client.Request(constants.Post, endPoint, payloadBytes)

	if err != nil {
		return nil, err
	}

	response, err := client.Client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	message := SlackMessage{}
	err = json.Unmarshal(body, &message)

	if err != nil {
		return nil, err
	}

	fmt.Printf("\nSuccess: send message to %+v\n\n", endPoint)

	return &message, nil
}

func (client *SlackClient) UpdateMessage(channel string, timestamp string, text string, callback *string) error {
	if channel == "" || timestamp == "" || text == "" {
		return exceptions.NewBadRequestError("Text and channel are required")
	}
	fmt.Print("\n\n============================\nStart update_message")

	payload := map[string]interface{}{
		"channel": channel,
		"text":    text,
		"ts":      timestamp,
	}

	if callback != nil {
		attachment := NewMessageAttachment(*callback)

		if _, err := json.Marshal(attachment); err != nil {
			return exceptions.NewBadRequestError("invalid attachments")
		} else {
			payload["attachments"] = *attachment
		}
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	endPoint := "/chat.update"
	request, err := client.Request(constants.Post, endPoint, payloadBytes)

	if err != nil {
		return err
	}

	response, err := client.Client.Do(request)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Printf("\nSuccess: update message to %+v\n\n", endPoint)

	return nil
}

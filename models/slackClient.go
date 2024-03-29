package models

import (
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

func (client *SlackClient) InitSlackClient() *SlackClient {
	httpClient := http.Client{
		Timeout: constants.RequestTimeOut * time.Second,
	}
	client.Client = &httpClient
	client.EndPoint = os.Getenv("SLACK_MESSAGE_API_ENDPOINT")

	return client
}

func (client *SlackClient) SlackRequest(method string, endpoint string, payload []byte) (*http.Request, error) {
	client = client.InitSlackClient()

	requestEndPoint := client.EndPoint + endpoint

	if request, err := http.NewRequest(method, requestEndPoint, bytes.NewBuffer(payload)); err != nil {
		return nil, err
	} else {
		// Set request headers
		additionHeaders := map[string]string{
			"Authorization": fmt.Sprintf("Bearer %+v", os.Getenv("SLACK_BOT_TOKEN")),
		}
		headers := GetHeaders(&additionHeaders)
		client.Headers = headers

		for key, value := range client.Headers {
			request.Header.Set(key, value)
		}

		return request, nil
	}
}

func (client *SlackClient) SlackConversationHistory(channel string, limit *int, timestamp *time.Time, inclusion *bool) (*SlackMessage, error) {
	if limit == nil {
		num := 20
		limit = &num
	}
	if timestamp == nil {
		now := time.Now()
		timestamp = &now
	}
	if inclusion == nil {
		defaultVal := false
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

	request, err := client.SlackRequest(constants.Post, "/conversations.history", payloadBytes)

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
	err = json.Unmarshal([]byte(string(body)), &message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func GetHeaders(additionHeaders *map[string]string) map[string]string {
	headers := make(map[string]string)

	defaultHeaders := map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
	}

	for key, value := range defaultHeaders {
		headers[key] = value
	}

	if additionHeaders != nil {
		for key, value := range *additionHeaders {
			headers[key] = value
		}
	}

	return headers
}

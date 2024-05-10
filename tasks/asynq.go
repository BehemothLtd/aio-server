package tasks

import (
	"os"

	"github.com/hibiken/asynq"
)

const (
	TypeDemoTask                          = "demo"
	SlackSendLeaveDayRequestMessageTask   = "slackSendLeaveDayRequest"
	SlackUpdateLeaveDayRequestStateTask   = "slackUpdteLeaveDayRequestState"
	SlackUpdateLeaveDayRequestMessageTask = "slackUpdateLeaveDayRequestMessage"
)

var AsynqClient *asynq.Client

func InitAsyncClient() *asynq.Client {
	if AsynqClient != nil {
		return AsynqClient
	}

	redisAddr := os.Getenv("REDIS_URL")

	AsynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	// defer AsynqClient.Close()

	return AsynqClient
}

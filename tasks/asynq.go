package tasks

import (
	"fmt"
	"os"

	"github.com/hibiken/asynq"
)

const (
	TypeDemoTask                        = "demo"
	SlackSendLeaveDayRequestMessagetask = "slackSendLeaveDayRequest"
)

var AsynqClient *asynq.Client

func InitAsyncClient() *asynq.Client {
	if AsynqClient != nil {
		return AsynqClient
	}

	redisAddr := os.Getenv("REDIS_URL")
	fmt.Printf("REDIS ADDR : %+v", redisAddr)

	AsynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	// defer AsynqClient.Close()

	return AsynqClient
}

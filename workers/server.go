package main

import (
	"aio-server/database"
	"aio-server/pkg/initializers"
	"aio-server/tasks"
	"log"
	"os"

	"github.com/hibiken/asynq"
)

func main() {
	initializers.LoadEnv()
	database.InitDb()

	worker := asynq.NewServer(
		asynq.RedisClientOpt{Addr: os.Getenv("REDIS_URL")},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6, // processed 60% of the time
				"default":  3, // processed 30% of the time
				"low":      1, // processed 10% of the time
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeDemoTask, tasks.HandleDemoTask)

	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}

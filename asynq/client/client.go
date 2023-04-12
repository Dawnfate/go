package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

// client.go
func main() {
	r := asynq.RedisClientOpt{Addr: "localhost:6379"}
	client := asynq.NewClient(r)

	// Create a task with typename and payload.
	t1 := asynq.NewTask(
		"email:welcome",
		map[string]interface{}{"user_id": 42})

	t2 := asynq.NewTask(
		"email:reminder",
		map[string]interface{}{"user_id": 42})

	// Process the task immediately.
	res, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %+v\n", res)

	// Process the task 24 hours later.
	res, err = client.EnqueueIn(24*time.Hour, t2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %+v\n", res)
}

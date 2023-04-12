package main

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

// ProcessTask should return nil if the task was processed successfully.
//
// If ProcessTask returns a non-nil error or panics, the task will be retried again later.
type Handler interface {
	ProcessTask(context.Context, *asynq.Task) error
}

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type {
	case "email:welcome":
		id, err := t.Payload.GetInt("user_id")
		if err != nil {
			return err
		}
		fmt.Printf("Send Welcome Email to User %d\n", id)

	case "email:reminder":
		id, err := t.Payload.GetInt("user_id")
		if err != nil {
			return err
		}
		fmt.Printf("Send Reminder Email to User %d\n", id)

	default:
		return fmt.Errorf("unexpected task type: %s", t.Type)
	}
	return nil
}

func main() {
	r := asynq.RedisClientOpt{Addr: "localhost:6379"}
	srv := asynq.NewServer(r, asynq.Config{
		Concurrency: 10,
	})

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}

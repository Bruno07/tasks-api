package main

import (
	"fmt"

	"github.com/Bruno07/tasks-api/internal/config"
	"github.com/Bruno07/tasks-api/internal/infra/queue"
	"github.com/Bruno07/tasks-api/internal/repositories"
)

func main() {

	config.LoadConfig()

	in := make(chan []byte)

	var repository = repositories.NewNotificationRepository(queue.GetInstanceQueue())

	msgs, _ := repository.Consumer("notification_queue", "notification")

	go func() {
		for m := range msgs {
			in <- []byte(m.Body)
		}
		close(in)
	}()

	for payload := range in {
		fmt.Println(string(payload))
	}

}
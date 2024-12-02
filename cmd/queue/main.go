package main

import (
	"fmt"

	"github.com/Bruno07/tasks-api/internal/config"
	"github.com/Bruno07/tasks-api/internal/infra/queue"
)

func main() {
	
    config.LoadConfig()

    in := make(chan []byte)
    ch := queue.GetInstanceQueue()
    queue.Consumer("notify_queue", "notify", ch, in)
    
    for payload := range in {
        fmt.Println(string(payload))
    }
}
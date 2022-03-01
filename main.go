package main

import (
	"context"
	"log"
	"math/rand"
	"time"

        "github.com/nrakochy/hello-wasm/tickler"
)

func main() {
}

//export enqueue
func Enqueue(y int) int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	service := NewService(ctx, 3)
	for i := 0; i < 10; i++ {
		if err := service.EnqueueRequest(Request(i)); err != nil {
			log.Fatalf("error sending request: %v", err)
			break
		}
		<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	for {
		time.Sleep(time.Second)
	}
}

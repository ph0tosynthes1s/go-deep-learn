package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

var blocker = make(chan struct{})

func routine(ctx context.Context, duration time.Duration) {
	defer close(blocker)
	select {
	case <-time.After(duration):
		fmt.Print("Yep!")
		blocker <- struct{}{}
	case <-ctx.Done():
		fmt.Printf("Operation timeout: %v\n", ctx.Err())
	}
}

func main() {
	randomCount := rand.IntN(5) + 1 // тип int
	randomDuration := time.Duration(randomCount) * time.Second
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go routine(ctx, randomDuration)
	<-blocker
}

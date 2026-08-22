package main

import (
	"context"
	"fmt"
)

func routine(ctx context.Context) {
	defer ctx.Done()
	val := ctx.Value("key")
	fmt.Println(val)
}

func main() {
	ctx := context.WithValue(context.Background(), "key", "value")
	routine(ctx)
}

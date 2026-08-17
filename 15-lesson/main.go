package main

import (
	"fmt"
	"time"
)

func check() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	go check()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Println(i)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}

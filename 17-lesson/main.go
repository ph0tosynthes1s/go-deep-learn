package main

import (
	"fmt"
	"math/rand"
)

var DATA = make(map[int]bool)
var signal chan struct{}

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

func writer(out chan<- int) {
	for {
		select {
		case <-signal:
			close(out)
			return

		case out <- random(1, 10):
		}
	}
}

func reader(in <-chan int, out chan<- int) {
	for i := range in {
		fmt.Print(i, " ")
		_, ok := DATA[i]
		if ok {
			signal <- struct{}{}
		} else {
			DATA[i] = true
			out <- i
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	A := make(chan int)
	B := make(chan int)

	signal = make(chan struct{})

	go writer(A)
	go reader(A, B)

	third(B)
}

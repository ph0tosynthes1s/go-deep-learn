package main

import (
	"fmt"
	"math/rand"
)

var CLOSEA bool
var DATA = make(map[int]bool)

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

func writer(out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(1, 10)
	}
}

func reader(in <-chan int, out chan<- int) {
	for i := range in {
		fmt.Print(i, " ")
		_, ok := DATA[i]
		if ok {
			CLOSEA = true
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
	go writer(A)
	go reader(A, B)
	third(B)
}

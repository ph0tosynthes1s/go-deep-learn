package main

import (
	"fmt"
	"sync"
)

func writeToChan(c chan int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- i
	close(c)
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int, 1)
	go writeToChan(c, 10, &wg)
	wg.Wait()
}

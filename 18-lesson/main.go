package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func channelRouter(wg *sync.WaitGroup, writeChan chan int, readChan chan int) {
	defer wg.Done()

	select {
	case write := <-writeChan:
		readChan <- write
		close(readChan)
	}
}

func reader(wg *sync.WaitGroup, in chan int) {
	defer wg.Done()

	fmt.Println(<-in)
}

func writer(wg *sync.WaitGroup, out chan int) {
	defer wg.Done()
	defer close(out)

	out <- rand.Int()
}

func main() {
	var wg sync.WaitGroup

	readChan := make(chan int)
	writeChan := make(chan int)

	wg.Add(3)

	go channelRouter(&wg, writeChan, readChan)
	go reader(&wg, readChan)
	go writer(&wg, writeChan)

	wg.Wait()
	fmt.Println("Программа успешно завершена")
}

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var myDate string

	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println("not enough")
	}

	// забираем дату из команды go run
	myDate = arguments[1]

	// парсим время в соответствии с константами го
	d, err := time.Parse("15:04 02 Jan 2006", myDate)

	if err == nil {
		fmt.Println(d)
	} else {
		fmt.Println("wrong formating")
	}
}

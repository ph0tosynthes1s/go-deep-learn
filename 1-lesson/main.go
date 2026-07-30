package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// число в формате инт64
	var sum int64

	// список аргументов при вызове go run
	arguments := os.Args

	// проверка на кол-во аргументов
	if len(arguments) == 1 {
		fmt.Println("Pls give more args!")
	}

	// цикл сборки аргументов
	for i := 0; i < len(arguments); i++ {
		num, err := strconv.ParseInt(arguments[i], 10, 64)
		if arguments[i] == "end" {
			break
		}
		if err != nil {
			fmt.Printf("%v is not a number!\n", arguments[i])
			continue
		}

		// складываем в объявленную ранее перменную
		sum += num
	}

	fmt.Println(sum)
}

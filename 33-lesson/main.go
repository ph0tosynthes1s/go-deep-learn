package main

import (
	"fmt"
	"os"
)

func main() {
	pathVal := os.Getenv("PATH")
	fmt.Println("PATH:", pathVal)

	val, ok := os.LookupEnv("MY_VAR")
	if ok {
		fmt.Println("MY_VAR =", val)
	} else {
		fmt.Println("Переменная MY_VAR не задана")
	}
}

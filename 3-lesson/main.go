package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println("not enough!")
		panic("not enough!")
	}

	// читаем файл из вызванной в терминале команды
	for _, filename := range arguments[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("err opening!")
			panic("err opening!")
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("err reading!")
				panic("err reading!")
			}
			fmt.Println(line)
		}
	}
}

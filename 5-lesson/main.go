package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Nums string
}

func main() {
	var myRecord Record

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("not enough")
	}

	filename := arguments[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open failed")
	}
	defer file.Close()

	decoderJSON := json.NewDecoder(file)

	err = decoderJSON.Decode(&myRecord)
	if err != nil {
		fmt.Println("parse failed")
	}
	fmt.Println(myRecord)

	myRecord = Record{
		Nums: " 5 4 3 2 1",
	}
	fmt.Println(myRecord)

	encocderJSON := json.NewEncoder(os.Stdout)
	err = encocderJSON.Encode(myRecord)
	if err != nil {
		fmt.Println("encode failed")
	}
}

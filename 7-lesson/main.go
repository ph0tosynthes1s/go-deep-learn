package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadXML(filename string, key interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("broken file!")
		return err
	}
	defer file.Close()

	decoderXML := xml.NewDecoder(file)
	err = decoderXML.Decode(key)
	return nil
}

func createJson(filename *os.File, key interface{}) error {
	encocderJSON := json.NewEncoder(filename)
	err := encocderJSON.Encode(key)
	if err != nil {
		fmt.Println("encode failed")
		return err
	}
	return nil
}

func main() {
	var xmlRecord Record

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("not enough args!")
		return
	}

	filename := arguments[1]

	err := loadXML(filename, &xmlRecord)
	if err == nil {
		fmt.Println(xmlRecord)
	}

	jsonFile, _ := os.Create("check.json")

	err = createJson(jsonFile, xmlRecord)
	if err == nil {
		fmt.Println("file created")
	}
}

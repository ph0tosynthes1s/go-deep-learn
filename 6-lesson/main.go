package main

import (
	"encoding/json"
	"fmt"
)

type Record struct {
	Name     string
	Surname  string
	Children []Child
}

type Child struct {
	Sex  string
	Name string
}

func main() {
	var newUnRec Record

	newRecord := Record{
		Name:    "Ivan",
		Surname: "Ivanov",
		Children: []Child{
			Child{
				Name: "Ivan",
				Sex:  "Male",
			},
			Child{
				Name: "Ivan",
				Sex:  "Male",
			},
			Child{
				Name: "Ivan",
				Sex:  "Male",
			},
		},
	}

	rec, err := json.Marshal(&newRecord)
	if err != nil {
		fmt.Println("troubles with marshall")
		return
	}
	fmt.Println(string(rec))

	err = json.Unmarshal(rec, &newUnRec)
	if err != nil {
		fmt.Println("troubles with unmarshall")
		return
	}
	fmt.Println(newUnRec)
}

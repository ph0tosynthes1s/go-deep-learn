package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 94
	var LENGTH int64 = 8

	arguments := os.Args

	switch len(arguments) {
	case 2:
		LENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
	default:
		fmt.Println("using default values!")
	}

	SEED := time.Now().Unix()
	rand.New(rand.NewSource(SEED))

	startChar := "!"

	var passwords []string

	for x := 0; x < 5; x++ {
		var i int64 = 1
		var passwordSlice []string
		for {
			myRand := random(MIN, MAX)
			newChar := string(startChar[0] + byte(myRand))
			if i == LENGTH {
				break
			}
			i++
			passwordSlice = append(passwordSlice, newChar)
		}
		password := strings.Join(passwordSlice, "")
		passwords = append(passwords, password)
	}
	now := time.Now().Hour()
	switch {
	case now < 12:
		fmt.Println(passwords[0])
	case now > 12 && now < 14:
		fmt.Println(passwords[1])
	case now > 14 && now < 16:
		fmt.Println(passwords[2])
	case now > 15 && now < 19:
		fmt.Println(passwords[3])
	}
}

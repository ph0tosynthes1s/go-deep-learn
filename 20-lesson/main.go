package main

import "math/rand"

func nPlus(n int) int {
	return n + 1 + 1
}

func main() {
	randNum := rand.Intn(5)
	nPlus(randNum)
}

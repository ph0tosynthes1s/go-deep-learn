package main

import "fmt"

func findOddNumber(arr []int) int {
	var result int
	for i := 0; i < len(arr); i++ {
		result = result ^ arr[i]
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 3, 1}
	fmt.Println(findOddNumber(arr))
}

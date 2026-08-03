package main

import "fmt"

// размер блока в хэш-мапе
const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashMap struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(value, size int) int {
	return (value % size)
}

func insert(hash *HashMap, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashMap) {
	for i := range hash.Table {
		if hash.Table[i] != nil {
			table := hash.Table[i]
			for table != nil {
				fmt.Printf("%d ->", table.Value)
				table = table.Next
			}
		}
		fmt.Println()
	}
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashMap{Table: table, Size: SIZE}
	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
}

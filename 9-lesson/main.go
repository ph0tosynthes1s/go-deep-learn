package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(node *Node, value int) int {
	// если корень пустой -> заполняем его
	if root == nil {
		node = &Node{value, nil}
		root = node
		return 0
	}

	if value == node.Value {
		fmt.Println("node already exists")
		return -1
	}

	if node.Next == nil {
		node.Next = &Node{value, nil}
		return -2
	}

	return addNode(node.Next, value)
}

func traverse(node *Node) {
	if node == nil {
		fmt.Println("empty nested list!")
	}

	for node != nil {
		fmt.Printf("%d ->", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func lookup(node *Node, value int) bool {
	if root == nil {
		node = &Node{value, nil}
		root = node
		return false
	}

	if value == node.Value {
		return true
	}

	if node.Next == nil {
		return false
	}

	return lookup(node.Next, value)
}

func getSize(node *Node) int {
	var nodeQuant int

	if node == nil {
		fmt.Println("nested list is empty")
		return 0
	}

	for node != nil {
		nodeQuant++
		node = node.Next
	}

	return nodeQuant
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("not enough args!")
		return
	}
	value, _ := strconv.Atoi(arguments[1])
	addNode(root, value)
	if lookup(root, value) {
		fmt.Println("node exists!")
	}
	addNode(root, 100)
	addNode(root, 200)
	addNode(root, 300)
	addNode(root, 400)
	addNode(root, 500)
	traverse(root)
	if getSize(root) > 0 {
		fmt.Println(getSize(root))
	}
}

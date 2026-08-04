package main

import "fmt"

type Node struct {
	Value float64
	Next  *Node
}

var size int = 0
var queue = new(Node)

func Push(node *Node, value float64) bool {
	if queue == nil {
		queue = &Node{value, nil}
		size++
		return true
	}

	node = &Node{value, nil}
	node.Next = queue

	queue = node

	size++
	return true
}

func traverse(node *Node) {
	if size == 0 {
		fmt.Println("queue is empty!")
	}

	for node != nil {
		fmt.Printf("%v->", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func Pop(node *Node) (float64, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return node.Value, true
	}

	temp := node
	for node.Next != nil {
		temp = node
		node = node.Next
	}

	value := (temp.Next).Value
	temp.Next = nil

	size--
	return value, true
}

func main() {
	queue = nil
	Push(queue, 1.23)
	Push(queue, 2.44)
	Push(queue, 2.46)
	Pop(queue)
	traverse(queue)
}

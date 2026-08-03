package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func create(value int) *Tree {
	var tree *Tree
	rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 2*value; i++ {
		temp := rand.Intn(value * 2)
		tree = insert(tree, temp)
	}
	return tree
}

func traverse(tree *Tree) {
	if tree == nil {
		return
	}
	traverse(tree.Left)
	fmt.Print(tree.Value, " ")
	traverse(tree.Right)
}

func insert(tree *Tree, value int) *Tree {
	if tree == nil {
		return &Tree{nil, value, nil}
	}
	if value == tree.Value {
		return tree
	}
	if value < tree.Value {
		tree.Left = insert(tree.Left, value)
		return tree
	}
	tree.Right = insert(tree.Right, value)
	return tree
}

func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is",
		tree.Value)
}

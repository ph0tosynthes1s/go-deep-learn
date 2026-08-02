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

func create(n int) *Tree {
	var tree *Tree
	rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
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

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
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

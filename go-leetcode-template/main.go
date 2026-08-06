package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tree = TreeNode{
	Val: 1,
	Left: &TreeNode{
		2,
		nil,
		&TreeNode{
			3,
			nil,
			nil}},
	Right: &TreeNode{
		2,
		nil,
		&TreeNode{
			3,
			nil,
			nil}}}

// [1,2,2,null,3,null,3]
func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode

	if root == nil {
		return true
	}

	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		l, r := queue[0], queue[1]
		queue = queue[2:]
		if l == nil && r == nil {
			continue
		} else if (l == nil && r != nil) || l != nil && r == nil {
			return false
		} else if l.Val != r.Val {
			return false
		}
		queue = append(queue, l.Left, r.Right, r.Left, l.Right)
	}
	return true
}
func main() {
	fmt.Println(isSymmetric(&tree))
}

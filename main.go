package main

import (
	"algorithms-go/binaryTree"
	"fmt"
)

func main() {
	tree := binaryTree.TreeNode{Val: 3, Left: &binaryTree.TreeNode{Val: 5}, Right: &binaryTree.TreeNode{Val: 10}}

	binaryTree.RecoverTree(&tree)

	fmt.Println(binaryTree.InorderTraversal(&tree))
}

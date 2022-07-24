package main

import (
	"algorithms-go/binaryTree"
	"algorithms-go/binaryTree/recoverBinarySearchTree"
	"fmt"
)

func main() {
	tree := binaryTree.TreeNode{
		Val: 1,
		Left: &binaryTree.TreeNode{Val: 3, Right: &binaryTree.TreeNode{Val: 2}},
	}

	recoverBinarySearchTree.RecoverTree(&tree)

	fmt.Println(binaryTree.InorderTraversal(&tree))
}

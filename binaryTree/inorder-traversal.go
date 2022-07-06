package binaryTree

//https://leetcode.com/problems/binary-tree-inorder-traversal/

func InorderTraversal(node *TreeNode) []int {
	var traversal []int
	var stack []*TreeNode

	for currentNode := node; len(stack) > 0 || currentNode != nil; {
		if currentNode == nil {
			currentNode = stack[len(stack)-1]
			stack = stack[0:(len(stack) - 1)]
		} else if currentNode.Left != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
			continue
		}

		traversal = append(traversal, currentNode.Val)
		currentNode = currentNode.Right
	}

	return traversal
}

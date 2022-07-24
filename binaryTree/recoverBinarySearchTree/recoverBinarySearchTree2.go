package recoverBinarySearchTree
import . "algorithms-go/binaryTree"

// The idea behind this algorithm is straight-forward.
// Given an inorder traversal of the given tree, we know that
// by definition, the swapped nodes are the first and last invalid
// nodes that are found. This solution does an iterative, inorder traversal,
// while always keeping track of the last 3 nodes process. We then use the last
// check the middle node of the last 3 to determine if it is valid or not. 
func RecoverTree(node *TreeNode)  {
	var stack []*TreeNode
	var nodes  [3]*TreeNode
	var firstInvalid, lastInvalid *TreeNode

	for currentNode := node; currentNode != nil || len(stack) > 0;  {
		if currentNode == nil {
			currentNode = stack[len(stack) - 1]
			stack = stack[0:(len(stack) - 1)]
		} else if currentNode.Left != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
			continue
		}

		nodes[0], nodes[1], nodes[2] = nodes[1], nodes[2], currentNode

		// process current node here
		if nodes[1] != nil {
			isLeftInvalid := nodes[0] != nil && nodes[1].Val < nodes[0].Val
			isRightInvalid := nodes[2] != nil && nodes[1].Val > nodes[2].Val

			if isRightInvalid || isLeftInvalid {
				if firstInvalid == nil {
					firstInvalid = nodes[1]
				} else {
					lastInvalid = nodes[1]
				}
			}
		}

		currentNode = currentNode.Right
	}

	if nodes[2].Val < nodes[1].Val {
		lastInvalid = nodes[2]
	}

	firstInvalid.Val, lastInvalid.Val = lastInvalid.Val, firstInvalid.Val
}


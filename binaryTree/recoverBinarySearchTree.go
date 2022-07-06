package binaryTree

//https://leetcode.com/problems/recover-binary-search-tree/
// The idea behind this algorithm exploits a couple of facts.
// First, given an inorder traversal we know that tree[i] <= tree[i + 1].
// Conversely, we know that tree[i] >= tree[i - 1]. That said, if we first
// iterate over the inorder traversal from left to right and check that tree[i] <= tree[i + 1],
// then the first node that violates the property is one of the nodes that was swapped.
// At the same time, if we iterate over the inorder traversal in reverse order and check that
// tree[i] >= tree[i + 1], then the first node that violates the property from that direction
// must be the second node that was swapped.

// The downsides to this particular solution are as follows:
// 1. Requires O(n) space
// 2. Requires traversing the tree 3 different times.
//    a) Yes, I know that I could've combined two of these loops
//       but doing so made the logic more complicated and doesn't
//       actually change the time complexity. So I opted for 2
//       different loops in favor of readability.

func RecoverTree(root *TreeNode)  {
	traversal := inorderTraversal(root)
	var swap1, swap2 *TreeNode

	for left := 0; left < len(traversal) - 1 && swap1 == nil; left++ {
		if traversal[left].Val > traversal[left + 1].Val {
			swap1 = traversal[left]
		}
	}

	for right := len(traversal) - 1; right > 0 && swap2 == nil; right-- {
		if traversal[right].Val < traversal[right - 1].Val {
			swap2 = traversal[right]
		}
	}

	swap1.Val, swap2.Val = swap2.Val, swap1.Val
}

func inorderTraversal(node *TreeNode) []*TreeNode {
	var traversal []*TreeNode
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

		traversal = append(traversal, currentNode)
		currentNode = currentNode.Right
	}

	return traversal
}

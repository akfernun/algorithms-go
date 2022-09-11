package linkedList
//https://leetcode.com/problems/reverse-linked-list/

//Definition for singly-linked list.
type ListNode struct {
  Val int
  Next *ListNode
}


func reverseList(head *ListNode) *ListNode {
	node := head
	var previousNode *ListNode

	for ; node != nil ; {
		nextNode := node.Next
		node.Next = previousNode

		previousNode = node
		node = nextNode
	}

	return previousNode
}


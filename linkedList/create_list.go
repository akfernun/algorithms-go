package linkedList

func CreateLinkedList(items []int) *ListNode {
	var prevNode *ListNode
	var head *ListNode

	for _, item := range items {
		node := &ListNode{Val: item}

		if head == nil {
			head = node
		}

		if prevNode != nil {
			prevNode.Next = node
		}

		prevNode = node
	}

	return head
}

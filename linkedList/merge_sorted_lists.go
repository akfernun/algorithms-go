package linkedList

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode

	// in this case, lets swap list1 and list2
	// to ensure that list1 always starts with the smallest
	// element. This means that we'll ALWAYS be merging INTO list1
	if list2.Val < list1.Val {
		head = list2
		list1, list2 = list2, list1
	} else {
		head = list1
	}

	for ;list1 != nil && list2 != nil; list1 = list1.Next {
		if list2.Val >= list1.Val && (list1.Next == nil || list2.Val <= list1.Next.Val) {
			temp := list2.Next
			list2.Next = list1.Next
			list1.Next = list2
			list2 = temp
		}
	}

	return head
}

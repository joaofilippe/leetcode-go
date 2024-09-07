package main


type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	for list1 != nil || list2 != nil{
		if list1.Val <= list2.Val{
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	return result.Next
}
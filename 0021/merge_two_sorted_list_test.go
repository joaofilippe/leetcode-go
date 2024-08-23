package mergeTwoSortedLists

import "testing"

func Test_MergeTwoSortedList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 9}
	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 6}

	mergedList := mergeTwoLists(l1, l2)
	println(mergedList.Val)
}

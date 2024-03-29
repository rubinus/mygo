package Solution3

import "mygo/data-str/05-Recursion/02-Test-Your-LinkedList-Solution/src/ListNode"

// 虚拟头结点
func RemoveElements3(head *ListNode.ListNode, val int) *ListNode.ListNode {
	dummyHead := &ListNode.ListNode{}
	dummyHead.Next = head

	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return dummyHead.Next
}

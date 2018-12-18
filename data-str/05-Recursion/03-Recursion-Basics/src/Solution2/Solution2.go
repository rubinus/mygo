package Solution2

import "mygo/data-str/05-Recursion/03-Recursion-Basics/src/ListNode"

func RemoveElements2(head *ListNode.ListNode, val int) *ListNode.ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return head
}

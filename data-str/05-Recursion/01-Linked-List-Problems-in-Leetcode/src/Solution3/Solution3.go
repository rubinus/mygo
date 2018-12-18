package Solution3

import "mygo/data-str/05-Recursion/01-Linked-List-Problems-in-Leetcode/src/ListNode"

/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/
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

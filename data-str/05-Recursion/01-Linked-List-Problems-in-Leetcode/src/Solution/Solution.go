package Solution

import "mygo/data-str/05-Recursion/01-Linked-List-Problems-in-Leetcode/src/ListNode"

/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/
func RemoveElements(head *ListNode.ListNode, val int) *ListNode.ListNode {
	for head != nil && head.Val == val {
		delNode := head
		head = head.Next
		delNode.Next = nil
	}
	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			cur := prev.Next
			prev.Next = cur.Next
			cur.Next = nil
		} else {
			prev = prev.Next
		}
	}

	return head
}

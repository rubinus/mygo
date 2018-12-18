package main

import (
	"fmt"
	"mygo/data-str/05-Recursion/02-Test-Your-LinkedList-Solution/src/ListNode"
	"mygo/data-str/05-Recursion/02-Test-Your-LinkedList-Solution/src/Solution3"
)

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := ListNode.GetListNode(nums)
	fmt.Println(head)

	res := Solution3.RemoveElements3(head, 6)
	fmt.Println(res)
}

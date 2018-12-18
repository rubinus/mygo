package main

import (
	"fmt"
	"mygo/data-str/05-Recursion/Optional-01-Recursive-LinkedList/src/LinkedList"
	"mygo/data-str/05-Recursion/Optional-01-Recursive-LinkedList/src/LinkedListR"
)

func main() {
	// 普通链表测试
	linkedList := LinkedList.GetLinkedList()
	for i := 0; i < 10; i++ {
		linkedList.AddLast(i)
	}
	linkedList.RemoveElement(0)
	linkedList.RemoveElement(6)
	linkedList.RemoveElement(9)
	fmt.Println(linkedList)

	// 递归实现的链表测试
	linkedListR := LinkedListR.GetLinkedListR()
	for i := 0; i < 10; i++ {
		linkedListR.AddFirst(i)
	}

	fmt.Println(linkedListR)
	for !linkedListR.IsEmpty() {
		fmt.Println("removed ", linkedListR.RemoveLast())
	}
	fmt.Println(linkedListR)
}

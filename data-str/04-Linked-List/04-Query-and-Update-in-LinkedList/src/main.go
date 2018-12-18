package main

import (
	"fmt"
	"mygo/data-str/04-Linked-List/04-Query-and-Update-in-LinkedList/src/LinkedList"
)

func main() {
	linkedList := LinkedList.GetLinkedList()

	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}

	linkedList.Add(2, 666)
	fmt.Println(linkedList)
}

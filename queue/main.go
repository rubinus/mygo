package main

import (
	"fmt"
	"mygo/queue/q"
)

func main() {
	q := queue.Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.IsEmpty()
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	q.Pop()
	q.IsEmpty()
}

package LinkedList

import "fmt"

type node struct {
	e    interface{}
	next *node
}

type LinkedList struct {
	head *node
	size int
}

// 获取链表中的元素个数
func (l *LinkedList) GetSize() int {
	return l.size
}

// 返回链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表头添加新的元素e
func (l *LinkedList) AddFirst(e interface{}) {
	//node := &node{e: e}
	//node.next = l.head
	//l.head = node
	l.head = &node{e, l.head}
	l.size++
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("Add failed.Illegal index.")
	}

	if index == 0 {
		l.AddFirst(e)
	} else {
		// 获得待插入节点的前一个节点
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// 插入新节点
		//node := &node{e: e, next: prev.next}
		//prev.next = node
		prev.next = &node{e, prev.next}
		l.size++
	}
}

// 在链表末尾添加新的元素e
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

func (n node) String() string {
	return fmt.Sprint(n.e)
}

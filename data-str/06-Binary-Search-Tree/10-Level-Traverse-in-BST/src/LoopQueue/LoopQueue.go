package LoopQueue

import (
	"bytes"
	"fmt"
)

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func GetLoopQueue(capacity int) (l *LoopQueue) {
	l = &LoopQueue{}
	l.data = make([]interface{}, capacity+1)
	l.front = 0
	l.tail = 0
	l.size = 0

	return
}

func (l *LoopQueue) GetCapacity() int {
	return len(l.data) - 1
}

func (l *LoopQueue) GetSize() int {
	return l.size
}

func (l *LoopQueue) IsEmpty() bool {
	return l.front == l.tail
}

// 入队
func (l *LoopQueue) Enqueue(e interface{}) {
	if (l.tail+1)%len(l.data) == l.front {
		l.resize(l.GetCapacity() * 2)
	}
	l.data[l.tail] = e
	l.tail = (l.tail + 1) % len(l.data)
	l.size++
}

// 获得队列头部元素
func (l *LoopQueue) Dequeue() (e interface{}) {
	if l.IsEmpty() {
		panic("cannot dequeue from empty queue")
	}

	e = l.data[l.front]
	l.data[l.front] = nil
	// 循环队列需要执行求余运算
	l.front = (l.front + 1) % len(l.data)
	l.size--
	if l.size == l.GetCapacity()/4 && l.GetCapacity()/2 != 0 {
		l.resize(l.GetCapacity() / 2)
	}

	return
}

// 查看队列头部元素
func (l *LoopQueue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("Queue is empty")
	}

	return l.data[l.front]
}

func (l *LoopQueue) resize(capacity int) {
	newData := make([]interface{}, capacity+1)
	for i := 0; i < l.size; i++ {
		newData[i] = l.data[(i+l.front)%len(l.data)]
	}
	l.data = newData
	l.front = 0
	l.tail = l.size
}

func (l *LoopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", l.size, l.GetCapacity()))
	buffer.WriteString("front [")
	for i := l.front; i != l.tail; i = (i + 1) % len(l.data) {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(l.data[i]))
		if (i+1)%len(l.data) != l.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}

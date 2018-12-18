package Array

import (
	"bytes"
	"fmt"
	"strconv"
)

type Array struct {
	// 声明类型为 slice
	data []int
	size int
}

// 传入数组的容量 capacity 返回 Slice
// 注：在 Go 中不同长度的数组属于不同类型，所以这里使用 Slice
func GetArray(capacity int) (a *Array) {
	a = &Array{}
	a.data = make([]int, capacity)
	a.size = 0
	return
}

// 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获得数组中的元素个数
func (a *Array) GetSize() int {
	return a.size
}

// 返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 向所有元素后添加一个新元素
func (a *Array) AddLast(element int) {
	//if a.size == len(a.data) {
	//	panic("AddLast failed,Array is full.")
	//}
	//
	//a.data[a.size] = element
	//a.size++
	a.Add(a.size, element)
}

// 向所有元素前添加一个新元素
func (a *Array) AddFirst(element int) {
	a.Add(0, element)
}

// 在第 index 个位置插入一个新元素 element
func (a *Array) Add(index int, element int) {
	if a.size == len(a.data) {
		panic("Add failed,Array is full")
	}

	if index < 0 || index > a.GetCapacity() {
		panic("Add failed,require index >= 0 and index <= a.cap")
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = element
	a.size++
}

// 获取 index 索引位置的元素
func (a *Array) Get(index int) int {
	if index < 0 || index >= a.size {
		panic("Get failed,Index is illegal.")
	}
	return a.data[index]
}

// 修改 index 索引位置的元素
func (a *Array) Set(index int, element int) {
	if index < 0 || index >= a.size {
		panic("Set failed,Index is illegal.")
	}
	a.data[index] = element
}

// 查找数组中是否有元素 element
func (a *Array) Contains(element int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return true
		}
	}

	return false
}

// 查找数组中元素 element 所在的索引，不存在则返回 -1
func (a *Array) Find(element int) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return i
		}
	}

	return -1
}

// 查找数组中元素 element 所有的索引组成的切片，不存在则返回 -1
func (a *Array) FindAll(element int) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			indexes = append(indexes, i)
		}
	}

	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (a *Array) Remove(index int) (element int) {
	if index < 0 || index >= a.size {
		panic("Set failed,Index is illegal.")
	}

	element = a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	return
}

// 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() int {
	return a.Remove(a.size - 1)
}

// 从数组中删除一个元素 element
func (a *Array) RemoveElement(element int) bool {
	index := a.Find(element)
	if index == -1 {
		return false
	}

	a.Remove(index)
	return true
}

// 从数组中删除所有元素 element
func (a *Array) RemoveAllElement(element int) bool {
	if a.Find(element) == -1 {
		return false
	}

	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			a.Remove(i)
		}
	}
	return true
}

// 重写 Array 的 string 方法
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(strconv.Itoa(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}

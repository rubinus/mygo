package main

import (
	"fmt"
	"mygo/data-str/08-Heap-and-Priority-Queue/05-Heapify-and-Replace-in-Heap/src/Array"
)

type MaxHeap struct {
	data *Array.Array
}

type freq struct {
	e         int
	frequency int
}

func GetMaxHeap() *MaxHeap {
	maxHeap := &MaxHeap{
		Array.GetArray(20),
	}

	return maxHeap
}

func (f *freq) compareTo(j *freq) int {
	if f.frequency > j.frequency {
		return -1
	} else if f.frequency < j.frequency {
		return 1
	} else {
		return 0
	}
}

// 返回堆中的元素个数
func (h *MaxHeap) Size() int {
	return h.data.GetSize()
}

// 返回一个布尔值, 表示堆中是否为空
func (h *MaxHeap) IsEmpty() bool {
	return h.data.IsEmpty()
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func (h *MaxHeap) parent(index int) int {
	if index == 0 {
		panic("index-0 doesn't have parent.")
	}
	return (index - 1) / 2
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (h *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (h *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

// 向堆中添加元素
func (h *MaxHeap) Add(e interface{}) {
	h.data.AddLast(e)
	h.siftUp(h.data.GetSize() - 1)
}

func (h *MaxHeap) siftUp(k int) {
	for k > 0 && h.data.Get(k).(*freq).compareTo(h.data.Get(h.parent(k)).(*freq)) > 0 {
		h.data.Swap(k, h.parent(k))
		k = h.parent(k)
	}
}

func (h *MaxHeap) FindMax() interface{} {
	if h.data.GetSize() == 0 {
		panic("Can not findMax when heap is empty.")
	}
	return h.data.Get(0)
}

func (h *MaxHeap) FindMin() interface{} {
	if h.data.GetSize() == 0 {
		panic("Can not findMax when heap is empty.")
	}
	return h.data.Get(h.data.GetSize() - 1)
}

// 取出堆中最大元素
func (h *MaxHeap) ExtractMax() interface{} {
	ret := h.FindMax()

	h.data.Swap(0, h.data.GetSize()-1)
	h.data.RemoveLast()
	h.siftDown(0)

	return ret
}

func (h *MaxHeap) ExtractMin() interface{} {
	ret := h.FindMin()

	h.data.RemoveLast()

	return ret
}

// data[j] 是 leftChild 和 rightChild 中的最大值
func (h *MaxHeap) siftDown(k int) {
	for h.leftChild(k) < h.data.GetSize() {
		j := h.leftChild(k)
		// j+1是右孩子索引，如果存在右孩子比较后获得左右孩子中较大值的索引
		if j+1 < h.data.GetSize() && h.data.Get(j+1).(*freq).compareTo(h.data.Get(j).(*freq)) > 0 {
			j++
		}
		// data[j] 是 leftChild 和 rightChild 中的最大值
		if h.data.Get(k).(*freq).compareTo(h.data.Get(j).(*freq)) > 0 {
			break
		}

		h.data.Swap(k, j)
		k = j
	}
}

// 取出堆中的最大元素，并且替换成元素e
func (h *MaxHeap) Replace(e interface{}) interface{} {
	ret := h.FindMax()

	h.data.Set(0, e)
	h.siftDown(0)

	return ret
}

/// 347. Top K Frequent Elements
/// https://leetcode.com/problems/top-k-frequent-elements/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的MaxHeap类
func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num]++
	}

	maxHeap := GetMaxHeap()
	for num, f := range numMap {
		if maxHeap.Size() < k {
			maxHeap.Add(&freq{e: num, frequency: f})
		} else {
			if f > maxHeap.FindMax().(*freq).frequency {
				maxHeap.Replace(&freq{e: num, frequency: f})
			}
		}
	}

	var res []int
	for !maxHeap.IsEmpty() {
		res = append(res, maxHeap.ExtractMax().(*freq).e)
	}
	return res
}

func main() {
	k := 2
	nums := []int{1, 1, 1, 2, 2, 3}

	fmt.Println(topKFrequent(nums, k))

	k = 1
	nums = []int{1}
	fmt.Println(topKFrequent(nums, k))
}

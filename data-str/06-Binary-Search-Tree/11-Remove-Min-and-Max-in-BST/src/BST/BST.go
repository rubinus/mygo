package BST

import (
	"bytes"
	"fmt"
	"strconv"
)

type node struct {
	e     int
	left  *node
	right *node
}

type BST struct {
	root *node
	size int
}

func GetBST() *BST {
	bst := &BST{}
	bst.root = nil
	bst.size = 0

	return bst
}

func (t *BST) GetSize() int {
	return t.size
}

func (t *BST) IsEmpty() bool {
	return t.size == 0
}

// 向二分搜索树中添加新的元素 e
func (t *BST) Add(e int) {
	t.root = t.add(t.root, e)
}

// 向以 node 为跟的二分搜索树中插入元素 e，递归算法
// 返回插入新节点后二分搜索树的根
func (t *BST) add(n *node, e int) *node {
	if n == nil {
		t.size++
		return &node{e: e}
	}

	// 递归调用
	if e < n.e {
		n.left = t.add(n.left, e)
	} else if e > n.e {
		n.right = t.add(n.right, e)
	}
	return n
}

// 看二分搜索树中是否包含元素 e
func (t *BST) Contains(e int) bool {
	return contains(t.root, e)
}

// 看以 node 为根的二分搜索树是否包含元素 e，递归算法
func contains(node *node, e int) bool {
	if node == nil {
		return false
	}

	if e == node.e {
		return true
	} else if e < node.e {
		return contains(node.left, e)
	} else {
		return contains(node.right, e)
	}
}

// 二分搜索树的前序遍历
func (t *BST) PreOrder() {
	preOrder(t.root)
}

// 前序遍历以 node 为根的二分搜索树，递归算法
func preOrder(node *node) {
	if node == nil {
		return
	}

	fmt.Println(node.e)
	preOrder(node.left)
	preOrder(node.right)
}

// 二分搜索树的非递归前序遍历
//func (t *BST) PreOrderNR() {
//	// 使用之前我们自己实现的数组栈
//	stack := GetArrayStack(20)
//	stack.Push(t.root)
//
//	for !stack.IsEmpty() {
//		cur := stack.Pop().(*node)
//		fmt.Println(cur.e)
//
//		if cur.right != nil {
//			stack.Push(cur.right)
//		}
//		if cur.left != nil {
//			stack.Push(cur.left)
//		}
//	}
//}

// 二分搜索树的中序遍历
func (t *BST) InOrder() {
	inOrder(t.root)
}

// 中序遍历以 node 为根的二分搜索树，递归算法
func inOrder(node *node) {
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Println(node.e)
	inOrder(node.right)
}

// 二分搜索树的后序遍历
func (t *BST) PostOrder() {
	postOrder(t.root)
}

// 后序遍历以 node 为根的二分搜索树，递归算法
func postOrder(node *node) {
	if node == nil {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.e)
}

// 二分搜索树的层序遍历
//func (t *BST) LevelOrder() {
//	// 使用我们之前实现的循环队列
//	queue := GetLoopQueue(20)
//	queue.Enqueue(t.root)
//	for !queue.IsEmpty() {
//		cur := queue.Dequeue().(*node)
//		fmt.Println(cur.e)
//
//		if cur.left != nil {
//			queue.Enqueue(cur.left)
//		}
//		if cur.right != nil {
//			queue.Enqueue(cur.right)
//		}
//	}
//}

// 寻找二分搜索树的最小元素
func (t *BST) Minimum() int {
	if t.size == 0 {
		panic("BST is empty!")
	}
	return minimum(t.root).e
}

// 返回以 node 为根的二分搜索树的最小值所在的节点
func minimum(n *node) *node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

// 寻找二分搜索树的最大元素
func (t *BST) Maximum() int {
	if t.size == 0 {
		panic("BST is empty!")
	}
	return maximum(t.root).e
}

// 返回以 node 为根的二分搜索树的最大值所在的节点
func maximum(n *node) *node {
	if n.right == nil {
		return n
	}
	return maximum(n.right)
}

// 从二分搜索树中删除最小值所在的节点，返回最小值
func (t *BST) RemoveMin() int {
	// 获得最小值
	ret := t.Minimum()
	t.root = t.removeMin(t.root)
	return ret
}

// 删除以 node 为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (t *BST) removeMin(n *node) *node {
	if n.left == nil {
		rightNode := n.right
		t.size--
		return rightNode
	}
	n.left = t.removeMin(n.left)
	return n
}

// 从二分搜索树中删除最小值所在的节点，返回最小值
func (t *BST) RemoveMax() int {
	// 获得最小值
	ret := t.Maximum()
	t.root = t.removeMax(t.root)
	return ret
}

// 删除以 node 为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (t *BST) removeMax(n *node) *node {
	if n.right == nil {
		leftNode := n.left
		t.size--
		return leftNode
	}
	n.right = t.removeMax(n.right)
	return n
}

func (t *BST) String() string {
	var buffer bytes.Buffer
	generateBSTSting(t.root, 0, &buffer)
	return buffer.String()
}

// 生成以 node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(node *node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + strconv.Itoa(node.e) + "\n")
	generateBSTSting(node.left, depth+1, buffer)
	generateBSTSting(node.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}

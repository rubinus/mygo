package mystd

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

/* 二叉树
1. 保存临时的左节点，然后交换左右节点
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//工厂函数
func CreateNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

//翻转
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := InvertTree(root.Left)
	root.Left = InvertTree(root.Right)
	root.Right = tmp
	return root
}

//根到叶子节点和
func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	currentVal := sum - root.Val
	return HasPathSum(root.Left, currentVal) ||
		HasPathSum(root.Right, currentVal)
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val + sumOfLeftLeaves(root.Right)
		} else {
			return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
		}
	} else {
		return sumOfLeftLeaves(root.Right)
	}
	return sum
}

/* 斐波那契数列
1：设置记忆搜索
*/
var fibMap = make(map[int]int)

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if fibMap[n] == 0 {
		fibMap[n] = Fib(n-1) + Fib(n-2)
	}
	return fibMap[n]
}

func Fib2(n int) int { //速度太慢不可取
	fib2 := make(map[int]int)
	fib2[n+1] = -1
	fib2[0] = 0
	fib2[1] = 1
	for i := 2; i <= n; i++ {
		fib2[i] = fib2[i-1] + fib2[i-2]
	}
	return fib2[n]
}

func StrReverse(str string) string {
	s := strings.Split(str, "")
	lenStr := len(s)
	for i := 0; i < lenStr/2; i++ {
		s[i], s[lenStr-1-i] = s[lenStr-1-i], s[i]
	}
	return strings.Join(s, "")
}

/* 零移动
1：给定一个数组 nums, 编写一个函数将所有 0 移动到它的末尾，同时保持非零元素的相对顺序。
2：例如， 定义 nums = [0, 1, 0, 3, 12]，调用函数之后， nums 应为 [1, 3, 12, 0, 0]。
3：记录k从0开始的位置，如果数组中nums[i]不等于0，确保k++，然后nums[k]和nums[i]交换位置，
4：当且仅当 i != k时，需要交换位置
*/
func MoveZero(nums []int) []int {
	k := 0
	for i, v := range nums {
		if v != 0 {
			if i != k {
				nums[k], nums[i] = v, 0
			}
			k++
		}
	}
	return nums
}

/* 2个数的和：滑动窗口自动右移
1：给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
2：你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
3：给定 nums = [2, 7, 11, 15], target = 9 返回 [0, 1]
4：利用map存取差值和对应的索引
5：查看是否存在map中，不存在的继续存，否则就取出对应的2个索引
6：把map中的索引放到返回值的第0位，是因为第一次执行时，map肯定是空的
7：当map第一次是空的时，索引是0的肯定先放进map
*/
func TwoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	res := make([]int, 2)
	for i, v := range nums {
		subInt := target - v
		if indexFromMap, ok := tmpMap[subInt]; ok {
			res[0] = indexFromMap //把map中的索引放到返回值的第0位，是因为第一次执行时，map肯定是空的
			res[1] = i
			return res
		}
		tmpMap[v] = i //当map第一次是空的时，索引是0的肯定先放进map
	}
	return nil
}

/* 判断一个数是否是素数
1：只有被1或是本身整除的数是互数
*/
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/* 最长子串主要特点：
1: 记录某个字符最后一次出现的位置到map中
2: 循环中的偏移量，逐渐右移
3: 长度随时++
4: 最长度==i+1减去偏移量
*/
func MaxLenStringNoRepSubstr(str string) int {
	lastPos := make(map[rune]int)
	start := 0
	maxLen := 0
	for i, ch := range []rune(str) {
		if lastIndex, ok := lastPos[ch]; ok && lastIndex >= start {
			start = lastIndex + 1 //偏移量右移
		}
		if i+1-start > maxLen { //记录最长的length
			maxLen = i + 1 - start
		}
		lastPos[ch] = i //记录最后出现的位置
	}
	return maxLen
}

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, j, max := 0, 1, 0
	for j < len(s) {
		x := strings.IndexByte(s[i:j], s[j])
		if x != -1 {
			max = Max(max, j-i)
			i = i + x + 1
		}
		j++
	}
	max = Max(max, j-i)
	return max
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v]++
		if len(m) == k+1 {
			delete(m, nums[i-k])
		}
	}
	return false
}

/* 二分查找主要特点: O(logN)
1: 必须在有序数组中取中间位置，循环不变量：[l...r]
2: 分别比较中间，左，右区间是否有这个数
3: 小于中间位置缩小左区间的右边界，否则缩小右区间的左边界
*/
func BinarySearch(arr []int, target int) int {
	lenArr := len(arr)
	l, r := 0, lenArr-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if target < arr[mid] { //搜索左区间
			r = mid - 1
		} else { //搜索右区间
			l = mid + 1
		}
	}
	return -1
}

func BinarySearchDG(arr []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if target == arr[mid] {
		return arr[mid]
	} else if target < arr[mid] {
		return BinarySearchDG(arr, l, mid-1, target)
	} else {
		return BinarySearchDG(arr, mid+1, r, target)
	}
}

func SortColors(nums []int) []int {
	zero, two := -1, len(nums)
	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i]

		} else {
			zero++
			nums[zero], nums[i] = nums[i], nums[zero]
			i++
		}
	}
	return nums
}

//长度最小的子数组
func MinSubArrayLen(s int, nums []int) int {
	l, r := 0, -1 //准备滑动窗口
	sum := 0
	res := len(nums) + 1
	for l < len(nums) {
		if r+1 < len(nums) && s > sum { //右边界右移扩大
			r++
			sum += nums[r]
		} else { //左边界右移缩小
			sum -= nums[l]
			l++
		}
		if sum >= s {
			res = int(math.Min((float64(res)), float64(r-l+1)))
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}
	return pre
}

func removeElements(head *ListNode, val int) *ListNode {
	var dummyHead = &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	if node.Next == nil {
		node = nil
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func binaryTreePaths(root *TreeNode) []string {
	ss := []string{}
	if root == nil {
		return ss
	}
	if root.Left == nil && root.Right == nil {
		ss = append(ss, strconv.Itoa(root.Val))
		return ss
	}
	ls := binaryTreePaths(root.Left)
	for _, v := range ls {
		ss = append(ss, strconv.Itoa(root.Val)+"->"+v)
	}
	rs := binaryTreePaths(root.Right)
	for _, v := range rs {
		ss = append(ss, strconv.Itoa(root.Val)+"->"+v)
	}
	return ss
}

func swapPairs(head *ListNode) *ListNode {
	var dummyHead = &ListNode{}
	dummyHead.Next = head
	p := dummyHead
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := node1.Next
		next := node2.Next

		p.Next = node2
		node2.Next = node1
		node1.Next = next

		p = node1
	}
	return dummyHead.Next
}

type DoubleNode struct {
	Value int
	// 前一个节点，以下统称为前指针
	PreNode *DoubleNode
	// 后一个节点，以下统称为后指针
	NextNode *DoubleNode
}

func ReverseDbNode(head *DoubleNode) *DoubleNode {
	// 先声明两个节点，并将两个节点都置为空
	var PreNode *DoubleNode
	var NextNode *DoubleNode
	// head为头节点，也就是当前节点
	for head != nil {
		// 保存第二个节点的值
		NextNode = head.NextNode
		// 将当前节点的前指针，指向后一个节点，也就是把头节点的后节点变成前节点，
		head.PreNode = NextNode
		// 将当前节点的后指针，指向前一个节点，也就是把头节点的前节点变成后节点。
		head.NextNode = PreNode
		// 更新前节点，也就是把头节点变成前节点
		PreNode = head
		// 更新当前节点，也就是把后节点变成头节点
		head = NextNode
	}
	return PreNode
}

/* 选择排序主要特点: O(n^2)，交换次数最少，稳定排序
1: 在每一轮中, 可以同时找到当前未处理元素的最大值和最小值
2: 迭代中每一次都查找数组中的最小值，把最小值放到已排序的后面，把最大值放到最右边的前面
3: 2轮迭代：逐个比较当前值与其它值的大小，碰撞式比较，从前向后，从后向前
4: 互换位置
5: 第一个for的 i 从0开始，第二个for的 j 从 i + 1开始
*/
func SelectSort(arr []int) {
	for i, lenArr := 0, len(arr); i < lenArr; i++ {
		for j := i + 1; j < lenArr; j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i] //位置互换
			}
		}
	}
}

/* 插入排序主要特点：O(n^2)  从索引1开始，copy一份，和前面的值比较，最后赋值
1: 从索引1开始，先保存一份当前index的值
2: 用这个值与左边的值进行比较，如果比左边的小，就把左边的值赋值给当前索引
3: 第一个for的 i 从 1 开始，第二个for逐个比较是否小于左边的元素
4: 最后把当前元素赋值给第二个for停止的位置
*/
func InsertSort(arr []int) {
	for i, lenArr := 1, len(arr); i < lenArr; i++ {
		index := arr[i] //取出一个位置比如i==3，和左边的元素比较
		var j int
		for j = i; j > 0 && index < arr[j-1]; j-- {
			arr[j] = arr[j-1] //把向左的值赋值给当前索引
		}
		arr[j] = index
	}
}

/* 归并排序主要特点：O(nlogN)  k表示将要放的数组索引，i是左区间，left是左边界，j是右区间，right是右边界，mid是中间
1: 数量<=100时优先使用insertSort
2: 把数组用递归的方法一分为二，分左边，分右边，再调用合并
3: 用临时数组保存合并时的数据，其长度是r - l + 1
4: 当merge所有的右区间时，临时数组的第一个元素下标都是减去l的偏移量
5: 左右区间比较，逐个放回原来的arr中
6: 左区间完了，直接放右区间；右区间完了直接放左区间
*/
func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	if r-l <= 100 { //小于100的数量改用插入排序
		InsertSort(arr)
		return
	}
	mid := l + (r-l)/2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	if arr[mid] > arr[mid+1] {
		tmpSlice := make([]int, r-l+1) //创建等同的slice
		for i := l; i <= r; i++ {
			tmpSlice[i-l] = arr[i]
		}
		merge(arr, tmpSlice, l, mid, r)
	}
}
func merge(arr, tmpSlice []int, l, mid, r int) {
	i, j := l, mid+1             //左右区间第一个元素开始设定初始值，然后比较
	lOffset, rOffset := i-l, j-l //关键一步，当右区间传过来的时候，第一个元素下标都是减去l的偏移量
	for k := l; k <= r; k++ {
		if tmpSlice[lOffset] < tmpSlice[rOffset] {
			arr[k] = tmpSlice[lOffset] //左右部分都有，开始比较，哪个小哪个放到arr中
			i++
		} else if tmpSlice[lOffset] > tmpSlice[rOffset] {
			arr[k] = tmpSlice[rOffset]
			j++
		} else if i > mid { //左边部分全部归并完了，直接把右边所有部分逐一合并
			arr[k] = tmpSlice[rOffset]
			j++
		} else if j > r { //右边部分全部归并完了，直接把左边所有部分逐一合并
			arr[k] = tmpSlice[lOffset]
			i++
		}
	}
}

/* 快速排序主要特点：为任意的基准数寻找位置 O(logN)
1: 数量<=100时优先使用insertSort
2: 随机取出任意一个基准数的值放到起始位置left，一直向右移动
3: 所有小于"基准"的元素，mid左边；[ left + 1 ... mid-1 ]
4: 所有大于"基准"的元素，mid右边。不用动
5: 把第一个位置放到基准数应该的位置 [ <v ] 自己就是基准数 [ >v ]
*/
func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	if r-l < 100 { //小于100的数量改用插入排序
		InsertSort(arr)
		return
	}
	pos := quick(arr, l, r) //pos是大于左区间，小于右区间的中间值
	QuickSort(arr, l, pos-1)
	QuickSort(arr, pos+1, r)
}

func quick(arr []int, l, r int) int {
	midV := arr[rand.Intn(r)]   //随机取出一个，避免接近排序的数组
	arr[l], midV = midV, arr[l] //和第一个元素交换位置
	//arr[l+1 ... mid-1] < v   arr[mid+1 ... i-1] > v
	mid := l //初始化最左边就是基准数的位置
	for i := l + 1; i <= r; i++ {
		if arr[i] < midV { //如果右移时比基准数的值小，就放到左边区间
			arr[mid+1], arr[i] = arr[i], arr[mid+1]
			mid++ //一直右移到自己应该的位置 [ <v ] 自己就是v [ >v ]
		}
	}
	arr[l], arr[mid] = arr[mid], arr[l]
	//循环完成后，把基准数换到应该在的位置
	return mid
}

/* 二路快速排序主要特点：为任意的基准数寻找位置 O(logN)
1: 数量<=100时优先使用insertSort
2: 随机取出任意一个基准数的值放到起始位置left，左边向右移，右边向左移
3: 所有小于"基准"的元素，leftMid左边；[ left + 1 ... leftMid)
4: 所有大于"基准"的元素，rightMid右边。(rightMid ... r]
5: 把第一个位置放到基准数应该的位置 [ <v ] 自己就是基准数 [ >v ]
*/
func QuickSort2(arr []int, l, r int) {
	if l >= r {
		return
	}
	if r-l < 100 { //小于100的数量改用插入排序
		InsertSort(arr)
		return
	}
	pos := quick2(arr, l, r) //pos是大于左区间，小于右区间的中间值
	QuickSort2(arr, l, pos-1)
	QuickSort2(arr, pos+1, r)
}
func quick2(arr []int, l, r int) int {
	midV := arr[rand.Intn(r)]   //随机取出一个，避免接近排序的数组
	arr[l], midV = midV, arr[l] //和第一个元素交换位置
	//arr[l+1 ... leftMid-1) <= v   arr(rightMid+1 ... i-1] >= v
	leftMid, rightMid := l+1, r //初始化最左边就是基准数的位置
	for leftMid < rightMid {    //在循环中，左边从左开始向右移，右边从右开始向左移
		for leftMid < rightMid && arr[leftMid] < midV && leftMid <= r {
			leftMid++
		}
		for leftMid < rightMid && arr[rightMid] > midV && rightMid >= l {
			rightMid--
		}
		arr[leftMid], arr[rightMid] = arr[rightMid], arr[leftMid] //交换位置
		leftMid++
		rightMid--
	}
	arr[l], arr[rightMid] = arr[rightMid], arr[l]
	//循环完成后，把基准数换到应该在的位置
	return rightMid
}

/* 三路快速排序主要特点：为任意的基准数寻找位置 O(logN)。l是基准位置，r是右边界，lt是小于基准位置++，gt是大于基准位置—，i自己++
1: 数量<=100时优先使用insertSort
2: 随机取出任意一个基准数的值放到起始位置left
3: 把所有数据的值arr[l]，分为 <  =  > 三部分，然后对 < 和 > 分别递归
4: 如果==的在中区间不动，< 的放到左区间，>的放到右区间
5: 互换基准元素到应该的位置
*/
func QuickSort3(arr []int, l, r int) {
	if l >= r {
		return
	}
	if r-l < 100 { //小于100的数量改用插入排序
		InsertSort(arr)
		return
	}
	lt, gt := quick3(arr, l, r) //pos是大于左区间，小于右区间的中间值
	QuickSort3(arr, l, lt-1)
	QuickSort3(arr, gt, r)
}
func quick3(arr []int, l, r int) (int, int) {
	midV := arr[rand.Intn(r)]   //随机取出一个，避免接近排序的数组
	arr[l], midV = midV, arr[l] //和第一个元素交换位置
	//arr[l+1 ... lt)   arr[lt+1 ... i)   arr[gt+1 ... r]
	i := l + 1
	lt, gt := l, r+1 //初始化最左边就是基准数的位置
	for i < gt {     //在循环中当快移动到右区间第一个大于基准数的值时
		if arr[i] < midV {
			arr[lt+1], arr[i] = arr[i], arr[lt+1]
			i++
			lt++
		} else if arr[i] > midV {
			arr[gt-1], arr[i] = arr[i], arr[gt-1]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	//循环完成后，把基准数换到应该在的位置
	return lt, gt
}

// 生成随机的数组，每个element不超过n的大小
// generate rand array
func CreateRandArr(n int) []int {
	testArr := make([]int, n)
	for k := range testArr {
		testArr[k] = rand.Intn(10000)
	}
	return testArr
}

package mystd

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// 查看测试代码覆盖率
// go test -coverprofile=c.out
// go tool cover -html=c.out

// go test -bench .
// go test -bench . -cpuprofile cpu.out
// go test -memprofile mem.out
// go tool pprof mem.out
// go tool pprof cpu.out   然后输入web  或是quit 保证下载了svg
// https://graphviz.gitlab.io/_pages/Download/Download_source.html

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
		i      int
	}{
		{
			[]int{2, 7, 9, 11},
			[]int{0, 1},
			9,
		},
		{
			[]int{20, 80, 90, 101},
			[]int{2, 3},
			191,
		},
	}

	for _, v := range tests {
		result := TwoSum(v.input, v.i)
		if !reflect.DeepEqual(result, v.output) {
			t.Error(result, v.output)
		} else {
			t.Log("yes", result, v.output)
		}
	}
}

func TestInvertTree(t *testing.T) {
	node := TreeNode{1, nil, nil}

	InvertTree(&node)
}

func TestHasPathSum(t *testing.T) {
	root := CreateNode(5)
	root.Left = CreateNode(4)
	root.Right = CreateNode(8)
	root.Left.Left = CreateNode(11)
	root.Left.Left.Left = CreateNode(7)
	root.Left.Left.Right = CreateNode(2)
	root.Right.Left = CreateNode(13)
	root.Right.Right = CreateNode(4)
	root.Right.Right.Right = CreateNode(1)

	HasPathSum(root, 22)
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(100000)
		//Fib2(4000)
	}
}
func TestFib(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{
			0,
			0,
		},
		{
			100000,
			2754320626097736315,
		},
		{
			10,
			55,
		},
		{
			6,
			8,
		},
	}

	for _, v := range tests {
		result := Fib(v.input)
		//result := Fib2(v.input)
		if result != v.output {
			t.Error(result, v.output)
		} else {
			t.Log("yes", result, v.output)
		}
	}
}

func TestBinarySearchDG(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			0,
		},
	}
	for _, v := range tests {
		i := BinarySearchDG(v.input, 0, len(v.input), v.output)
		fmt.Println(i)
	}
}

func TestStrReverse(t *testing.T) {
	s := StrReverse("ABC")
	fmt.Println(s)

	ss := []string{"a", "z", "w"}
	sort.Strings(ss)
	fmt.Println(ss)
}

func TestSelectSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		//{
		//	[]int{10, 4, 6, 2, 3, 3, 8, 0, 1, 100, 90},
		//	[]int{0, 1, 2, 3, 3, 4, 6, 8, 10, 90, 100},
		//},
		{
			[]int{5, 2, 8, 4, 9, 1},
			[]int{1, 2, 4, 5, 8, 9},
		},
		//{
		//	[]int{10},
		//	[]int{10},
		//},
		//{
		//	[]int{},
		//	[]int{},
		//},
		//{
		//	[]int{0,0,0,0},
		//	[]int{0,0,0,0},
		//},
		//{
		//	[]int{0,1},
		//	[]int{0,1},
		//},
	}

	for _, v := range tests {
		SelectSort(v.input)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Error(v.input, v.output)
		} else {
			t.Log("yes", v.input, v.output)
		}
	}
}
func BenchmarkSelectSort(b *testing.B) {
	arr := CreateRandArr(10000)
	for i := 0; i < b.N; i++ {
		SelectSort(arr)
		//fmt.Println(arr)
	}
}
func ExampleSelectSort() {
	arr := []int{10, 4, 6, 2, 3}
	SelectSort(arr)
	//Output
	//[2,3,4,6,10]
}

func TestInsertSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			[]int{10, 4, 6, 2, 3, 8, 0, 1, 100, 90},
			[]int{0, 1, 2, 3, 4, 6, 8, 10, 90, 100},
		},
		{
			[]int{10},
			[]int{10},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{0, 4, 3, 1},
			[]int{0, 1, 3, 4},
		},
	}

	for _, v := range tests {
		InsertSort(v.input)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Error(v.input, v.output)
		} else {
			t.Log("yes", v.input, v.output)
		}
	}
}
func BenchmarkInsertSort(b *testing.B) {
	arr := CreateRandArr(100)
	for i := 0; i < b.N; i++ {
		InsertSort(arr)
		//fmt.Println(arr)
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			[]int{10, 0, 0, 0, 0, 4, 6, 2, 3, 8, 0, 1, 100, 90, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 6, 8, 10, 90, 100},
		},
		{
			[]int{10},
			[]int{10},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{0, 1},
			[]int{0, 1},
		},
	}

	for _, v := range tests {
		MergeSort(v.input, 0, len(v.input)-1)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Error(v.input, v.output)
		} else {
			t.Log("yes", v.input, v.output)
		}
	}
}
func BenchmarkMergeSort(b *testing.B) {
	arr := CreateRandArr(100000)
	for i := 0; i < b.N; i++ {
		MergeSort(arr, 0, len(arr)-1)
		//fmt.Println(arr)
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			[]int{10, 4, 0, 0, 0, 6, 2, 3, 8, 0, 1, 100, 90},
			[]int{0, 0, 0, 0, 1, 2, 3, 4, 6, 8, 10, 90, 100},
		},
		{
			[]int{10},
			[]int{10},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{0, 1},
			[]int{0, 1},
		},
	}

	for _, v := range tests {
		QuickSort3(v.input, 0, len(v.input)-1)
		if !reflect.DeepEqual(v.input, v.output) {
			t.Error(v.input, v.output)
		} else {
			t.Log("yes", v.input, v.output)
		}
	}
}
func BenchmarkQuickSort(b *testing.B) {
	arr := CreateRandArr(10000)
	for i := 0; i < b.N; i++ {
		QuickSort(arr, 0, len(arr)-1)
		//fmt.Println(arr)
	}
}

func BenchmarkQuickSort2(b *testing.B) {
	arr := CreateRandArr(10000)
	for i := 0; i < b.N; i++ {
		QuickSort2(arr, 0, len(arr)-1)
		//fmt.Println(arr)
	}
}

func BenchmarkQuickSort3(b *testing.B) {
	arr := CreateRandArr(10000)
	for i := 0; i < b.N; i++ {
		QuickSort3(arr, 0, len(arr)-1)
		//fmt.Println(arr)
	}
}

func TestQuickSort3(t *testing.T) {
	arr := []int{10, 4, 6, 2, 3}
	QuickSort3(arr, 0, len(arr)-1)
}

func ExampleQuickSort3() {
	arr := []int{10, 4, 6, 2, 3}
	QuickSort3(arr, 0, len(arr)-1)
	//fmt.Println(arr)
	//Output
	//[2,3,4,6,10]
}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"mygo/data-str/08-Heap-and-Priority-Queue/04-Extract-and-Sift-Down-in-Heap/src/MaxHeap"
	"time"
)

func main() {
	n := 100000

	maxHeap := MaxHeap.GetMaxHeap()
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		randNum := rand.Intn(math.MaxInt32)
		maxHeap.Add(randNum)
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = maxHeap.ExtractMax().(int)
	}
	fmt.Println(arr)

	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			panic("Error")
		}
	}

	fmt.Println("Test MaxHeap completed.")
}

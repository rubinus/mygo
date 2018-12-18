package main

import (
	"fmt"
	"mygo/data-str/02-Arrays/05-Contain-Find-and-Remove/src/Array"
)

func main() {
	arr := Array.GetArray(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)
	// [-1, 0, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9]

	arr.Remove(2)
	fmt.Println(arr)

	arr.RemoveElement(4)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)

	arr.AddLast(1)
	indexes := arr.FindAll(1)
	fmt.Println(arr, indexes)

	isRemove := arr.RemoveAllElement(1)
	fmt.Println(arr, isRemove)
}

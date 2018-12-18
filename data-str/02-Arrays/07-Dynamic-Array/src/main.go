package main

import (
	"fmt"
	"mygo/data-str/02-Arrays/07-Dynamic-Array/src/Array"
)

func main() {
	arr := Array.GetArray(10)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)
}

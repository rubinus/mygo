package main

import (
	"fmt"
	"mygo/data-str/02-Arrays/03-Add-Element-in-Array/src/Array"
)

func main() {
	a := Array.GetArray(5)

	a.AddLast(2)
	a.AddLast(3)
	a.AddFirst(1)
	fmt.Println(a)
}

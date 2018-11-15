package main

import (
	"fmt"
)

func fn(str string) chan rune {
	ch := make(chan rune)
	go func() {
		for _, s := range []rune(str) {
			ch <- s
		}
		close(ch)
	}()
	return ch
}
func main() {
	str := "你日1024节happy不"
	var result string
	for r := range fn(str) {
		result += string(r)
	}
	fmt.Println(result)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice1 := a[5:10]
	fmt.Println(slice1)
	//[6,7,8,9,10]
	a[5] = 10
	fmt.Println(slice1)
	//[10,7,8,9,10]
	slice1 = append(slice1, 100)
	fmt.Println(slice1)
	//[10,7,8,9,10,100]
	a[6] = 11
	fmt.Println(slice1)
	fmt.Println(a)
	//[10,7,8,9,10,100]

}

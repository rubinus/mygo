package main

import "fmt"

/*
@Time : 2019-07-25 10:31
@Author : rubinus.chu
@File : main
@project: mygo
*/

func main() {
	var s = make([]int, 2)
	//var s []int

	s = append(s, 1)
	fmt.Println(cap(s))

	s = append(s, 2)
	fmt.Println(cap(s))
	s = append(s, 3)
	fmt.Println(cap(s))

	r := s[1:2]

	r2 := r[2:3]

	fmt.Println(s)

	fmt.Println(r)

	fmt.Println(r2)
}

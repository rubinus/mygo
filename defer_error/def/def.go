package main

import "fmt"

func def() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
func main() {
	def()
}

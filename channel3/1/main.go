package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 20)

	go func() {
		for v := range ch {
			fmt.Println(v, "--v--")
		}
	}()

	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	time.Sleep(1 * time.Second)
	for i := 10; i < 20; i++ {
		ch <- i
	}

	time.Sleep(5 * time.Second)
}

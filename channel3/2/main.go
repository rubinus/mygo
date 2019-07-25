package main

import "fmt"

var (
	numCh  = make(chan int)
	quitCh = make(chan chan bool)
)

func main() {
	go handler()

	for i := 0; i < 3; i++ {
		numCh <- i
	}
	wait := make(chan bool)
	quitCh <- wait

	fmt.Println("done!")
}

func handler() {
	for {
		select {
		case res := <-numCh:
			fmt.Println(res)
		case <-quitCh:
			fmt.Println("done-123")
			return
		}
	}
}

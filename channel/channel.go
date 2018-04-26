package main

import (
	"fmt"
	"sync"
	"strconv"
	"time"
)

type workers struct {
	in   chan string
	done func()
}

//channel做为参数
func doWork(i int, work workers) {
	go func() {
		//fmt.Println("beging...")
		for {
			n := <-work.in
			fmt.Println(i, n)
			work.done()
		}
	}()
}

//channel做为返回值
func createWorker(i int, group *sync.WaitGroup) workers {
	work := workers{
		in: make(chan string),
		done: func() {
			group.Done()
		},
	}
	go doWork(i, work)
	return work
}
func chanDemo() {
	works := [10]workers{}
	var group sync.WaitGroup
	for i, _ := range works {
		works[i] = createWorker(i, &group)
	}
	group.Add(20)
	for i, w := range works {
		w.in <- "A" + strconv.Itoa(i)
	}
	for i, w := range works {
		w.in <- "B" + strconv.Itoa(i)
	}
}

func generate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * 100)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	chanDemo()

	c1, c2 := generate(), generate()
	tick := time.Tick(time.Second)
	var c1s []int
	for {
		active := make(chan int)
		var ac int
		if len(c1s) > 0 {
			ac = c1s[0]
		}
		select {
		case n := <-c1:
			c1s = append(c1s, n)
			fmt.Println("c1 from ..", n)
		case n := <-c2:
			c1s = append(c1s, n)
			fmt.Println("c2 from...", n)
		case active <- ac:
			fmt.Println("slice has")
			c1s = c1s[1:]
		case <-tick:
			fmt.Println("c1 = ", len(c1s))
		default:
			fmt.Println("no c1 and c2")
			time.Sleep(time.Second)
		}
	}

}

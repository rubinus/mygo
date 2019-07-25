package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	F102()
}

func F102() {
	defer func() { fmt.Println(1) }()

	defer func() { fmt.Println(2) }()

	defer func() { fmt.Println(3) }()
}

// -----
func test2() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // a=1,b=1	先算第二个calc函数
	// 1 1 3 4
	a = 0
	defer calc("2", a, calc("20", a, b)) // a=0,b=2
	// 2 0 2 2
	b = 1
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// ----------
func test() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//------------
type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println(m)

}

// -------
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

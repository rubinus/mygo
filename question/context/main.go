package main

import (
	"context"
	"fmt"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	// 定义 goroutine 函数
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 0
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("ctx.Done() is run")
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // main 方法执行完后，结束 ctx 相关的 goroutine

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

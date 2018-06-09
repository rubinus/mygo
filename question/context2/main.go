package main

import (
	"context"
	"fmt"
	"time"
)

func Cdd(ctx context.Context) int {
	fmt.Println(ctx.Value("NLJB"))
	fmt.Println("Cdd key1", ctx.Value("HELLO"))
	time.Sleep(3 * time.Second)

	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return -3
	default:
		// 没有结束 ... 执行 ...
		return 300
	}
}
func Bdd(ctx context.Context) int {
	fmt.Println("key1", ctx.Value("HELLO"))
	fmt.Println("key2", ctx.Value("WROLD"))
	ctx = context.WithValue(ctx, "NLJB", "NULIJIABEI")
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	go fmt.Println(Cdd(ctx))

	time.Sleep(2 * time.Second)

	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return -2
	default:
		// 没有结束 ... 执行 ...
		return 200

	}
}
func Add(ctx context.Context) int {
	ctx = context.WithValue(ctx, "HELLO", "WROLD")
	ctx = context.WithValue(ctx, "WROLD", "HELLO")
	go fmt.Println(Bdd(ctx))

	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return -1
	default:
		// 没有结束 ... 执行 ...
		return 100

	}
}
func main() {
	// 自动取消(定时取消)
	{
		timeout := 1 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		fmt.Println(Add(ctx))
	}
	// 手动取消
	//  {
	//      ctx, cancel := context.WithCancel(context.Background())
	//      go func() {
	//          time.Sleep(2 * time.Second)
	//          cancel() // 在调用处主动取消
	//      }()
	//      fmt.Println(Add(ctx))
	//  }
}

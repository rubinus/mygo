package backend

import (
	"context"
	"fmt"
	"testing"

	"code.tvmining.com/tvplay/tvmq/config"
)

func TestSendFilter(t *testing.T) {
	body := FilterBody{
		Items: []Items{
			{
				Id:     "1001",
				UserId: "123",
				Text:   "我们",
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()
	in := SendFilter(ctx, config.FilterHost, body)
	select {
	case <-ctx.Done():
		fmt.Println("超时")
	case result := <-in:
		fmt.Println(result)

	}
}

func BenchmarkSendFilter(b *testing.B) {

	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()
	body := FilterBody{
		Items: []Items{
			{
				Id:     "1001",
				UserId: "123",
				Text:   "我们",
			},
		},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		in := SendFilter(ctx, config.FilterHost, body)
		select {
		case <-ctx.Done():
			fmt.Println("超时")
		case result := <-in:
			fmt.Println(result)

		}
	}
}

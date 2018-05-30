package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	out := 50 * time.Millisecond
	d := time.Now().Add(out)
	fmt.Println(d)
	//ctx, cancel := context.WithDeadline(context.Background(), d)
	ctx, cancel := context.WithTimeout(context.Background(), out)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
		fmt.Println(ctx.Deadline())

	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println(ctx.Deadline())
	}
}

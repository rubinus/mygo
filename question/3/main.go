package main

import (
	"fmt"
	"math"
)

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", "123")
}
func reverse(x int) int {
	num := 0
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	// 使用 math 包中定义好的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return num
}

const (
	x = iota
	y
	z = "z"
	k
	p = iota
)

func main() {
	c := &ConfigOne{}
	c.String()

	fmt.Println(reverse(-123))
	fmt.Println(x, y, z, k, p)
}

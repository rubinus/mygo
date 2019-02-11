package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(luckCode())
}
func luckCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(10000)
	return code
}

package main

import (
	"fmt"
	"math/rand"
	"mygo/data-str/11-Union-Find/06-Path-Compression/src/UF"
	"mygo/data-str/11-Union-Find/06-Path-Compression/src/UnionFind3"
	"mygo/data-str/11-Union-Find/06-Path-Compression/src/UnionFind4"
	"mygo/data-str/11-Union-Find/06-Path-Compression/src/UnionFind5"
	"time"
)

func testUF(uf UF.UF, m int) time.Duration {
	size := uf.GetSize()
	rand.Seed(time.Now().Unix())

	startTime := time.Now()

	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.UnionElements(a, b)
	}

	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.IsConnected(a, b)
	}

	return time.Now().Sub(startTime)
}

func main() {
	// UnionFind1 慢于 UnionFind2
	//size := 100000
	//m := 10000

	// UnionFind2 慢于 UnionFind1, 但UnionFind3最快
	size := 100000
	m := 100000

	//uf1 := UnionFind1.Constructor(size)
	//fmt.Println(testUF(uf1, m))
	//
	//uf2 := UnionFind2.Constructor(size)
	//fmt.Println(testUF(uf2, m))

	uf3 := UnionFind3.Constructor(size)
	fmt.Println(testUF(uf3, m))

	uf4 := UnionFind4.Constructor(size)
	fmt.Println(testUF(uf4, m))

	uf5 := UnionFind5.Constructor(size)
	fmt.Println(testUF(uf5, m))
}

package main

import (
	"fmt"
	"mygo/data-str/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/AVLSet"
	"mygo/data-str/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/BSTSet"
	"mygo/data-str/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/FileOperation"
	"mygo/data-str/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/LinkedListSet"
	"mygo/data-str/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/Set"
	"path/filepath"
	"time"
)

func testSet(set Set.Set, filename string) time.Duration {
	startTime := time.Now()

	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words:", len(words))
	for _, word := range words {
		set.Add(word)
	}
	fmt.Println("Total different words:", set.GetSize())

	return time.Now().Sub(startTime)
}

func main() {
	filename, _ := filepath.Abs("data-str/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/pride-and-prejudice.txt")

	bstSet := BSTSet.Constructor()
	time1 := testSet(bstSet, filename)
	fmt.Println("BST Set :", time1)

	linkedListSet := LinkedListSet.Constructor()
	time2 := testSet(linkedListSet, filename)
	fmt.Println("linked List Set:", time2)

	avlSet := AVLSet.Constructor()
	time3 := testSet(avlSet, filename)
	fmt.Println("AVL Set:", time3)
}

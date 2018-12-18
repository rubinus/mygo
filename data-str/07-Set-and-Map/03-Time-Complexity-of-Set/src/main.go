package main

import (
	"fmt"
	"mygo/data-str/07-Set-and-Map/03-Time-Complexity-of-Set/src/BSTSet"
	"mygo/data-str/07-Set-and-Map/03-Time-Complexity-of-Set/src/FileOperation"
	"mygo/data-str/07-Set-and-Map/03-Time-Complexity-of-Set/src/LinkedListSet"
	"mygo/data-str/07-Set-and-Map/03-Time-Complexity-of-Set/src/Set"
	"os"
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
	projectPath, _ := os.Getwd()
	currentPath := filepath.Join(projectPath, "07-Set-and-Map", "03-Time-Complexity-of-Set")

	filename := filepath.Join(currentPath, "pride-and-prejudice.txt")

	bstSet := BSTSet.GetBSTSet()
	time1 := testSet(bstSet, filename)
	fmt.Println("BST set :", time1)

	linkedListSet := LinkedListSet.GetLinkedListSet()
	time2 := testSet(linkedListSet, filename)
	fmt.Println("linkedList set:", time2)
}

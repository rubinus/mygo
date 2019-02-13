package main

import (
	"bufio"
	"fmt"
	"mygo/pipeline/pkg"
	"os"
)

func main() {
	arr := []int{3, 2, 6, 7, 4}
	ch1 := pkg.ArraySource(arr)

	arr2 := []int{7, 4, 0, 3, 2, 13, 8}
	ch2 := pkg.ArraySource(arr2)

	in := pkg.Merge(ch1, ch2)

	for v := range in {
		fmt.Println(v)
	}

	filename := "small.in"
	//写
	file, e := os.Create(filename)
	if e != nil {
		panic(e)

	}
	defer file.Close()
	source := pkg.RandomSource(5000000)
	writer := bufio.NewWriter(file)
	pkg.WriterSink(writer, source)
	writer.Flush()
	fmt.Println("write done")

	//读
	open, e := os.Open(filename)
	if e != nil {
		panic(e)
	}
	defer open.Close()
	readerSource := pkg.ReaderSource(bufio.NewReader(open), -1)
	count := 0
	for v := range readerSource {
		count++
		fmt.Println(v)
		if count >= 10 {
			break
		}
	}
	fmt.Println(count)

}

package main

import (
	"bufio"
	"fmt"
	"mygo/pipeline/pkg"
	"os"
)

func main() {
	ch := createPipeline("small.in", 5000000, 4)
	writerFile("small.out", ch)
	readFile("small.out")
}

func readFile(s string) {

	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	ch := pkg.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range ch {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}

}

func writerFile(filename string, ints chan int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	defer buf.Flush()

	pkg.WriterSink(buf, ints)

}

func createPipeline(filename string, filesize, chunkcount int) chan int {
	chunks := filesize / chunkcount
	var sortChan []chan int
	for i := 0; i < chunkcount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		//读到什么地方
		file.Seek(int64(chunks*i), 0)
		ch := pkg.ReaderSource(bufio.NewReader(file), chunks)
		sortChan = append(sortChan, pkg.MemInSort(ch))

	}
	return pkg.MergeN(sortChan...)
}

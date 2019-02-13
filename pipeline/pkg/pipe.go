package pkg

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func ArraySource(a []int) chan int {
	out := make(chan int)
	go func() {
		sort.Ints(a)
		for _, v := range a {
			out <- v

		}
		close(out)
	}()
	return out
}
func MemInSort(ch chan int) chan int {
	out := make(chan int, 1024)
	go func() {
		var arr []int
		for v := range ch {
			arr = append(arr, v)
		}
		fmt.Println("read done:", time.Now().Sub(startTime))
		sort.Ints(arr)
		fmt.Println("sort done:", time.Now().Sub(startTime))
		for _, v := range arr {
			out <- v
		}
		close(out)
	}()
	return out

}

func Merge(in1, in2 chan int) chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}

		}
		close(out)
		fmt.Println("merge done:", time.Now().Sub(startTime))
	}()
	return out
}

func MergeN(input ...chan int) chan int {
	if len(input) == 1 {
		return input[0]
	}
	m := len(input) / 2
	return Merge(MergeN(input[:m]...), MergeN(input[m:]...))

}

func ReaderSource(reader io.Reader, chunkSize int) chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		readSize := 0
		for {
			n, err := reader.Read(buffer)
			readSize += n

			if n > 0 {
				u := int(binary.BigEndian.Uint64(buffer))
				out <- u
			}
			if err != nil || (chunkSize != -1 && readSize >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(writer io.Writer, in chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(c int) chan int {
	out := make(chan int, 1024)
	go func() {
		for i := 0; i < c; i++ {
			out <- rand.New(rand.NewSource(time.Now().UnixNano())).Int()
		}
		close(out)
	}()
	return out
}

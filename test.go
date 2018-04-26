package main

import (
	"./scrypt"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	dk, _ := scrypt.Key([]byte("s"), []byte("ABCD"), 32768, 8, 1, 32)

	privateKey := base64.StdEncoding.EncodeToString(dk)

	fmt.Println(privateKey)

	h := sha256.New()
	io.WriteString(h, privateKey)

	fmt.Println(h.Sum(nil))

	f, err := os.Open("test.go")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(errors.New("没有这个文件"))
	}

	data := []int{1, 2, 3}

	for i, v := range data {

		v *= 10 //original item is not changed

		data[i] = v
	}

	abc := 123

	fmt.Println(123, abc)
	fmt.Println(123)
	fmt.Println("data:", data) //prints data: [1 2 3]

	if f != nil {
		log.Fatal(f, "----f-----")
	}

}

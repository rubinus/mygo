package main

import (
	"mygo/morerequest/do"
)

//var tvmid = "sjh5a7ab7ac84143d74c92b29be"
//var wx_token = "33580c57d3c86f07"

func main() {

	c := make(chan int)
	for i := 0; i < 10; i++ {
		go do.DoWork(i)
	}
	<-c

}

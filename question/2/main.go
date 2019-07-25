package main

import (
	"fmt"
	"time"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

	F101()

	time.Sleep(time.Second)

}

func F101() {
	sl := []string{"one", "two", "three"}
	for _, v := range sl {
		//v := v
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second)
}

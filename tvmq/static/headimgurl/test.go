package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/models/user"
	"code.tvmining.com/tvplay/tvmq/utils"
)

func main() {
	f, err := os.Create("./static/headimgurl/newUsers.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ch := make(chan string)
	saveUser(ch)
	for v := range ch {
		fmt.Println(v)
		s := strings.Split(v, " ")

		u := models.ResComment{
			Userid:     s[0],
			Headimgurl: "https://" + config.SocketHost + "/headimgurl/photos/" + s[1] + ".jpg",
			Nickname:   s[2],
		}
		b, _ := json.Marshal(u)
		f.WriteString(string(b) + "\n")
	}

}

func saveUser(ch chan string) {
	//f, err := os.Open("./users.txt")
	f, err := os.Open("./static/headimgurl/users.txt")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(f)

	go func() {
		defer f.Close()

		for {
			b, _, err := buf.ReadLine()
			if err != nil {
				if err == io.EOF {
					fmt.Println("读完了")
					break
				}
				panic(err)
			}
			result := string(b)
			//fmt.Println(result)
			s := strings.Split(result, " ")

			//fmt.Println(s[0],s[1])
			//"http://qa-sc.tvplay.tvm.cn/headimgurl/photos/"
			u := &user.User{
				Headimgurl: s[0] + ".jpg",
				Nickname:   s[1],
				Country:    "中国",
				IsAi:       "1",
				MinAppid:   config.Minappid,
				MinOpenid:  strconv.Itoa(int(utils.GetCurrentTime(19))),
			}
			uid, err := user.SaveUser(u)
			if err != nil {
				fmt.Println("save is failed")
			}
			//fmt.Println(uid)

			str := fmt.Sprintf("%s %s %s", uid, s[0], s[1])
			ch <- str
		}
	}()

}

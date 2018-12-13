package handler

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/utils"
)

type HostConn struct {
	Delay time.Duration
	Key   string
}

func (hc HostConn) Run() {
	go func() {

		//ip := utils.GetIntranetIp()
		ip, _ := os.Hostname()
		lib.JudgeSadd(hc.Key, ip)
		fmt.Println(ip, "提供服务")

		for range time.Tick(hc.Delay) {
			CheckOnlineMac()
		}

	}()
}

func CheckOnlineMac() {
	reply, err := lib.JudgeSmembers(config.OnlineHostConnKey)
	if err != nil {
		fmt.Sprintf("Smembers is failed %s", err.Error())
		return
	}

	c := make(chan string)
	go func() {
		for _, v := range reply {
			c <- v
		}
		close(c)
	}()

	for host := range c {
		var errStr string

		go func(host string) {

			s := utils.ParseDns(host)
			if s == false {
				errStr = host + "超时"
				reply, err := lib.JudgeSrem(config.OnlineHostConnKey, host)
				if err != nil {
					fmt.Printf("连接删除%s失败%s%s", host, reply, err.Error())
				} else {
					fmt.Println(errStr + "--已删除")
				}
			}

			//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			//defer cancel()
			//out := RunCMD(ctx, host)
			//select {
			//case <-ctx.Done():
			//	errStr = host + "超时"
			//	reply, err := lib.JudgeSrem(config.OnlineHostConnKey, host)
			//	if err != nil {
			//		fmt.Printf("连接删除%s失败%s%s", host, reply, err.Error())
			//	} else {
			//		fmt.Println(errStr + "--已删除")
			//	}
			//case <-out:
			//	//fmt.Println(host, " is ok")
			//}
		}(host)

	}
}

func RunCMD(ctx context.Context, command string) chan string {
	out := make(chan string)
	go func() {
		_, err := exec.Command("ping", "-c", "3", command).Output()
		if err != nil {
			out <- ""
		} else {
			out <- "ok"
		}
	}()
	return out
}

func RunCMD2(ctx context.Context, command string) chan string {
	out := make(chan string)
	go func() {
		in := bytes.NewBuffer(nil)
		cmd := exec.Command("ping", "-c", "1", command)
		cmd.Stdin = in
		in.WriteString(command + "\n")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			out <- ""
		}
		defer stdout.Close()
		if err := cmd.Start(); err != nil {
			out <- ""
		}
		opBytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			out <- ""
		}
		out <- string(opBytes)
	}()
	return out

}

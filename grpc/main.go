package main

import (
	"fmt"
	"mygo/grpc/example"

	"github.com/golang/protobuf/proto"
)

func main() {
	// 创建一个消息 Info
	info := &example.Helloworld{
		Id:  345,
		Str: "678",
	}
	// 进行编码
	data, err := proto.Marshal(info)
	if err != nil {
		fmt.Printf("marshaling error: ", err)
	}

	// 进行解码
	newInfo := &example.Helloworld{}
	err = proto.Unmarshal(data, newInfo)
	if err != nil {
		fmt.Printf("unmarshaling error: ", err)
	}

	fmt.Println(info.GetId(), "---", newInfo.GetId())
	if info.GetId() != newInfo.GetId() {
		fmt.Printf("data mismatch %q != %q", info.GetId(), newInfo.GetId())
	} else {
		fmt.Printf("%+v", newInfo)
	}
}

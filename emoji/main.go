package main

import (
	"fmt"
	"regexp"
)

func main() {
	var emojiRx = regexp.MustCompile(`[\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	var s = emojiRx.ReplaceAllString("🍎初雨繁花 Thats a nice joke 😆😆😆 😛", `[e]`)
	fmt.Println(s)
}

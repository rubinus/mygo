package main

import (
	"fmt"
	"math"
	"regexp"
	"unicode"
)

func main() {
	num := GetStrLength("，中国人t!")
	fmt.Println(len([]rune("，中国人t!")))
	fmt.Println(num)
}

// GetStrLength 返回输入的字符串的字数，汉字和中文标点算 1 个字数，英文和其他字符 2 个算 1 个字数，不足 1 个算 1个
func GetStrLength(str string) float64 {
	var total float64

	reg := regexp.MustCompile("/·|，|。|《|》|‘|’|”|“|；|：|【|】|？|（|）|、/")

	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || reg.Match([]byte(string(r))) {
			total = total + 1
		} else {
			total = total + 0.5
		}
	}

	return math.Ceil(total)
}

package fetch

import (
	"bufio"
	"io"
	"net/http"

	"errors"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetcher(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("错误的地址:" + url)
	}
	defer resp.Body.Close()

	e := GetCharset(resp.Body)
	utf8body := transform.NewReader(resp.Body, e.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(utf8body)
	if err != nil {
		return nil, errors.New("goQuery解析错误:")
	}
	return doc, nil
}

func GetCharset(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

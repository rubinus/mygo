package parser

import (
	"fmt"
	"mygo/crawler/fetcher"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/transform"
)

func TestParseCity(t *testing.T) {
	resp, err := http.Get("http://www.zhenai.com/zhenghun/anshan")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	e := fetch.GetCharset(resp.Body)
	utf8body := transform.NewReader(resp.Body, e.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(utf8body)
	if err != nil {
		panic(err)
	}
	doc.Find("#ResultLbox .content a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		city := s.Text()
		fmt.Println(link, city)
	})
}

func TestParsePerson(t *testing.T) {
	doc, err := fetch.Fetcher("http://album.zhenai.com/u/108037476")
	if err != nil {
		panic(err)
	}
	ParsePerson(doc)

}

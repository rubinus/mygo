package parser

import (
	"mygo/crawler/engine"

	"mygo/crawler/model"

	"os"

	"fmt"

	"strconv"
	"strings"

	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/json-iterator/go"
)

func ParseCityList(noInter interface{}) engine.ParseResult {
	result := engine.ParseResult{}
	d := noInter.(*goquery.Document)
	d.Find("#cityList a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		city := s.Text()
		result.Items = append(result.Items, city)
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        link,
				ReqFlag:    "City",
				ParserFunc: ParseCity,
			})
	})
	return result
}

func ParseCity(noInter interface{}) engine.ParseResult {
	result := engine.ParseResult{}
	d := noInter.(*goquery.Document)
	d.Find("#ResultLbox .content a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		name := s.Text()
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        link,
				ReqFlag:    "Person",
				ParserFunc: ParsePerson,
			})
	})
	return result
}

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d+])岁</td>`)

func extStr(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return ""
	} else {
		return ""
	}

}
func ParsePerson(noInter interface{}) engine.ParseResult {
	result := engine.ParseResult{}
	d := noInter.(*goquery.Document)

	person := model.Person{}
	d.Find(".brief-table td").Each(func(i int, s *goquery.Selection) {
		valueAll := s.Text()
		value := strings.Split(valueAll, "：")[1]
		fmt.Println(value, "--person--", i)

		switch i {
		case 0:
			ageNum := strings.Replace(value, "岁", "", 1)
			age, _ := strconv.Atoi(ageNum)
			person.Age = age
		case 1:
			heiNum := strings.Replace(value, "CM", "", 1)
			height, _ := strconv.Atoi(heiNum)
			person.Height = height
		case 3:
			person.Marriage = value
		case 4:
			person.Education = value
		case 8:
			person.Hukou = value
		}

	})

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := jsonIterator.Marshal(person) //json_iterator
	os.Stdout.Write(b)

	return result
}

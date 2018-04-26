package engine

import "mygo/crawler/fetcher"

func worker(r Request) (ParseResult, error) {
	body, err := fetch.Fetcher(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}

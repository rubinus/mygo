package engine

import (
	"fmt"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(rs ...Request) {
	var q []Request

	for _, v := range rs {
		q = append(q, v)
	}

	for len(q) > 0 { //如果队列中有，就取第一个
		seed := q[0]
		q = q[1:]
		fmt.Println(seed.Url)

		results, err := worker(seed)
		if err != nil {
			panic(err)
		}

		q = append(q, results.Requests...) //continue to queue

		//if seed.ReqFlag == "CityList" && len(results.Items) > 1 && len(results.Requests) > 1 {
		//	q = append(q, results.Requests[0:10]...) //continue to queue
		//} else if seed.ReqFlag == "City" {
		//	q = append(q, results.Requests...) //continue to queue
		//}

	}

}

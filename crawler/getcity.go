package crawler

import (
	"mygo/crawler/engine"
	"mygo/crawler/scheduler"
	"mygo/crawler/zhenai/parser"
)

func GetCity() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ReqFlag:    "CityList",
	//	ParserFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ReqFlag:    "CityList",
		ParserFunc: parser.ParseCityList,
	})

}

package engine

type Request struct {
	Url        string
	ReqFlag    string
	ParserFunc func(interface{}) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser(noInter interface{}) ParseResult {
	return ParseResult{}
}

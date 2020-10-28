package engine

type Request struct {
	Url string
	ParserFunc func([] byte) ParseResult
}

//解析结果
type ParseResult struct {
	Reuqests []Request
	Items []interface{}
}

//暂时还没用到
func NilParser([]byte) ParseResult  {
	return ParseResult{}
}
package engine

type Request struct {
	Url string
	ParserFunc func([] byte) ParseResult
}

type ParseResult struct {
	Reuqests []Request
	Items []interface{}
}

func NilParser([]byte) ParseResult  {
	return ParseResult{}
}
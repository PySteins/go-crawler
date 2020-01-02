package engine

type Request struct {
	Url        string
	Item       Item
	ParserFunc func(string, Request) ParseResult
}

type ParseResult struct {
	Requests []Request
	//Items    []Item
}

func NilParser(string, Request) ParseResult {
	return ParseResult{}
}

type Item struct {
	Name    string
	Chapter string
	Images  []string
}

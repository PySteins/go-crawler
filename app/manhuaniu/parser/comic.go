package parser

import (
	"crawler/engine"
	"regexp"
)

var comicRe = regexp.MustCompile(`<p class="ell"><a href="(https://www.manhuaniu.com/manhua/[1-9]+/)">([^<]+)</a></p>`)

func ComicParser(body string, r engine.Request) engine.ParseResult {
	rs := engine.ParseResult{}
	match := comicRe.FindAllStringSubmatch(body, -1)
	for _, m := range match {
		r.Item.Name = m[2]
		r.Url = m[1]
		r.ParserFunc = ChapterParser
		rs.Requests = append(rs.Requests, r)
	}
	return rs
}

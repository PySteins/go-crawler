package parser

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

var chapterRe = regexp.MustCompile(`<a href="(/manhua/[1-9]+/[1-9]+\.html)"[^<]+<span>([^<]+)</span>`)

func ChapterParser(body string, r engine.Request) engine.ParseResult {
	rs := engine.ParseResult{}

	match := chapterRe.FindAllStringSubmatch(body, -1)
	for _, m := range match {
		r.Item.Chapter = m[2]
		r.Url = fmt.Sprintf("https://www.manhuaniu.com%s", m[1])
		r.ParserFunc = ImageParser
		rs.Requests = append(rs.Requests, r)
	}
	return rs
}

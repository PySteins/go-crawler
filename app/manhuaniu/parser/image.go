package parser

import (
	"crawler/engine"
	"fmt"
	"regexp"
	"strings"
)

var imageRe = regexp.MustCompile(`"(images[^"]+)"`)

func ImageParser(body string, r engine.Request) engine.ParseResult {
	rs := engine.ParseResult{}
	match := imageRe.FindAllStringSubmatch(body, -1)
	var imgList []string
	for _, m := range match {
		url := fmt.Sprintf("https://restp.dongqiniqin.com//%s", m[1])
		url = strings.Replace(url, `\`, "", -1)
		imgList = append(imgList, url)
	}
	r.Item.Images = imgList
	r.ParserFunc = engine.NilParser
	r.Url = ""
	rs.Requests = append(rs.Requests, r)
	return rs
}

package engine

import (
	"crawler/utils"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ParseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, ParseResult.Requests...)

		for _, request := range ParseResult.Requests {
			log.Printf("Got item %v", request.Item)
		}
	}
}

func Worker(r Request) (ParseResult, error) {
	// log.Printf("fetching %s", r.Url)
	body, err := utils.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher error url: %s, error: %s", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r), nil
}

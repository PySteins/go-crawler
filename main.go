package main

import (
	"crawler/app/manhuaniu/parser"
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		ItemChan: persist.ItemSaver("comic", "manhuaniu"),
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "https://www.manhuaniu.com/rank/",
		ParserFunc: parser.ComicParser,
		//Url:        "https://www.manhuaniu.com/manhua/256/",
		//ParserFunc: parser.ChapterParser,
		//Url:        "https://www.manhuaniu.com/manhua/256/3366.html",
		//ParserFunc: parser.ImageParser,
	})
}

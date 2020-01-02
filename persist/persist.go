package persist

import (
	"context"
	"crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index, types string) chan engine.Item {
	ch := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Println("open client err: ", err)
		return nil
	}

	go func() {
		for item := range ch {
			err = save(client, index, types, item)
			if err != nil {
				log.Println("Item save err: ", err)
			}
		}
	}()
	return ch
}


func save(client *elastic.Client, index, types string, item engine.Item) error {
	_, err := client.Index().
		Index(index).
		Type(types).
		BodyJson(item).Do(context.Background())
	return err
}
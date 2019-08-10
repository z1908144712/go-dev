package main

import (
	"context"

	"github.com/olivere/elastic"
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}

func initES() (err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		return
	}
	tweet := LogMessage{
		App:     "olivere",
		Message: "Take five",
	}
	_, err = client.Index().Index("twitter").Type("Tweet").Id("1").BodyJson(tweet).Do(context.Background())
	return
}

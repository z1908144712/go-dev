package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println(err)
		return
	}
	tweet := Tweet{
		User:    "olivere",
		Message: "Take five",
	}
	_, err = client.Index().Index("twitter").Type("Tweet").Id("1").BodyJson(tweet).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}

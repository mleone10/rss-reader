package main

import (
	"log"

	"github.com/mleone10/rss-reader/pkg/dynamo"
	"github.com/mleone10/rss-reader/pkg/rss"
)

func main() {
	r := rss.NewRssReader(&dynamo.HardcodedFeedLister{})
	feeds, err := r.ProcessFeeds()
	if err != nil {
		log.Fatalf("everything is broken: %v", err)
	}

	for _, f := range feeds {
		log.Println(f.Channel.Items)
	}
}

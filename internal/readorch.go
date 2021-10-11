package internal

import (
	"fmt"
	"log"
)

type ReadOrchestrator struct {
	rssReader rssReader
	db        *dynamoDataStore
}

func (o ReadOrchestrator) ProcessFeeds() error {
	if o.db == nil {
		o.db, _ = newDynamodbClient()
	}

	feedUrls, err := o.db.listFeeds()
	if err != nil {
		return fmt.Errorf("Failed to list rss feeds: %v", err)
	}

	for _, f := range o.rssReader.ReadAll(feedUrls) {
		o.processFeed(f)
	}

	return nil
}

func (o ReadOrchestrator) processFeed(f Feed) {
	for _, i := range f.Channel.Items {
		log.Println(i.Title)
	}
}

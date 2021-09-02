package internal

import (
	"fmt"
	"log"

	"github.com/mleone10/rss-reader/pkg/rss"
)

type FeedLister interface {
	ListAllFeeds() ([]string, error)
}

func ProcessFeeds(f FeedLister) error {
	r := rss.NewRssReader()

	feedUrls, err := f.ListAllFeeds()
	if err != nil {
		return fmt.Errorf("Failed to list rss feeds: %v", err)
	}

	fs := r.ReadAll(feedUrls)
	for _, f := range fs {
		for _, i := range f.Channel.Items {
			log.Println(i.Title)
		}
	}

	return nil
}

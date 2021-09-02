package rss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RssReader struct {
	feedLister FeedLister
	httpClient http.Client
}

type FeedLister interface {
	ListAll() ([]string, error)
}

func NewRssReader(feedLister FeedLister) RssReader {
	return RssReader{
		feedLister: feedLister,
		httpClient: http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (r *RssReader) ProcessFeeds() ([]Feed, error) {
	feedUrls, err := r.feedLister.ListAll()
	if err != nil {
		return nil, fmt.Errorf("Failed to list rss feeds: %v", err)
	}

	return r.readAll(feedUrls), nil
}

func (r *RssReader) readAll(feedUrls []string) []Feed {
	feeds := []Feed{}
	for _, f := range feedUrls {
		feed, err := r.read(f)
		if err != nil {
			// TODO: Log error with individual stream
		} else {
			feeds = append(feeds, *feed)
		}
	}
	return feeds
}

func (r *RssReader) read(feedUrl string) (*Feed, error) {
	res, err := r.httpClient.Get(feedUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get rss feed: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var feed Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal channel response to xml: %v", err)
	}

	return &feed, nil
}

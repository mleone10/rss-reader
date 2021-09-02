package rss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RssReader struct {
	httpClient http.Client
}

func NewRssReader() RssReader {
	return RssReader{
		httpClient: http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (r *RssReader) ReadAll(feedUrls []string) []Feed {
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

package internal

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type rssReader struct {
	httpClient *http.Client
}

func (r rssReader) ReadAll(feedUrls []string) []Feed {
	if r.httpClient == nil {
		r.httpClient = &http.Client{
			Timeout: 5 * time.Second,
		}
	}

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

func (r rssReader) read(feedUrl string) (*Feed, error) {
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

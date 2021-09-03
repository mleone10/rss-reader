package main

import (
	"github.com/mleone10/rss-reader/internal"
	"github.com/mleone10/rss-reader/pkg/dynamo"
)

func main() {
	internal.ProcessFeeds(dynamo.HardcodedFeedLister{})
}

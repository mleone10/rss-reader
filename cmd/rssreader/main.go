package main

import (
	"github.com/mleone10/rss-reader/internal"
)

func main() {
	internal.ReadOrchestrator{}.ProcessFeeds()
}

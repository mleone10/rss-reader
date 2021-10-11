package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mleone10/rssreader/internal"
)

func main() {
	lambda.Start(internal.ReadOrchestrator{}.ProcessFeeds)
}

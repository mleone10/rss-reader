package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mleone10/rss-reader/internal"
	"github.com/mleone10/rss-reader/pkg/dynamo"
)

func HandleRequest(ctx context.Context) error {
	return internal.ProcessFeeds(dynamo.HardcodedFeedLister{})
}

func main() {
	lambda.Start(HandleRequest)
}

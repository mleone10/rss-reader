package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoDataStore struct {
	client *dynamodb.DynamoDB
}

func newDynamodbClient() (*dynamoDataStore, error) {
	session := session.Must(session.NewSession())
	client := dynamodb.New(session, aws.NewConfig().WithRegion("us-east-1"))

	return &dynamoDataStore{
		client: client,
	}, nil
}

func (d dynamoDataStore) listFeeds() ([]string, error) {
	return []string{
		"https://marioleone.me/index.xml",
		"https://www.guildwars2.com/en/feed/",
	}, nil
}

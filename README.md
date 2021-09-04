# RSS-Reader

This RSS Reader microservice is designed to execute multiple times per day. On each execution, it reads a list of RSS feed URLs, fetches each feed and stores new entries in a database. It then publishes events for downstream services to listen and react to.

Currently, the service is deployed to an AWS Lambda Function and uses a DynamoDB Table for both the list of RSS feed URLs and the record of read entries. New entries are published to AWS SNS.

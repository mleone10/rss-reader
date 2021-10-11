terraform {
  backend "s3" {
    bucket = "leone-terraform-states"
    key    = "rss-reader.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_iam_role" "lambda_role" {
  name = "${var.project_name}-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
  inline_policy {
    name = "${var.project_name}-policy"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          "Effect" : "Allow",
          "Action" : "dynamodb:Query"
          "Resource" : "*"
        }
      ]
    })
  }
}

resource "aws_lambda_function" "lambda_function" {

  function_name = "${var.project_name}-lambda"
  filename      = "../handler.zip"
  role          = aws_iam_role.lambda_role.arn
  handler       = "bin/rssreaderlambda"
  runtime       = "go1.x"
}

resource "aws_sns_topic" "sns_topic" {
  name = "${var.project_name}-new-events"
}

resource "aws_dynamodb_table" "dynamodb_table" {
  name           = "${var.project_name}-items"
  billing_mode   = "PROVISIONED"
  write_capacity = 1
  read_capacity  = 1
  hash_key       = "channel"
  range_key      = "item_guid"

  attribute {
    name = "channel"
    type = "S"
  }

  attribute {
    name = "item_guid"
    type = "S"
  }

  attribute {
    name = "channel_update_ts"
    type = "S"
  }

  global_secondary_index {
    name           = "channels"
    write_capacity = 1
    read_capacity  = 1
    hash_key       = "channel"
    range_key      = "channel_update_ts"
  }
}

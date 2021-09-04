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

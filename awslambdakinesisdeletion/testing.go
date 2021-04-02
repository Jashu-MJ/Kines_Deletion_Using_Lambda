package main

import (
	"awslambdakinesisdeletion/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(handler.Handler)
}

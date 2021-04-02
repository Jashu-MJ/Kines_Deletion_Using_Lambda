package main

import (
	"awslambdakinesiscreation/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(handler.Handler)
}

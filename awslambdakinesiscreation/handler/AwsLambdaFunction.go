//Package handler has lambda function for kinesis stream creation.
package handler

import (
	"awslambdakinesiscreation/api"
	"awslambdakinesiscreation/awskinesis"

	"github.com/aws/aws-lambda-go/events"
)

//Handler is the starting point of Lambda function, It executes from here to
// create a kinesis stream for appropriate request recieved from Api gateway
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apiBody, resp, err := api.ValidatingTheAPIRequest(req) // validating the api request
	if err != nil {
		return resp, nil
	}

	kinesisClient, resp, err := awskinesis.KinesisClientCreation() // kinesis client creation if valid request.
	if err != nil {
		return api.Resp444, nil
	}

	resp = awskinesis.CreateKinesisStream(apiBody, kinesisClient) // kinesis stream creation
	return resp, nil
}

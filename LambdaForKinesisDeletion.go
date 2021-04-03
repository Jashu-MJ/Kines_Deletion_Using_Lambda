//Package Handler has lambda function used for kinesis deletion.
package handler

import (
	"awslambdakinesisdeletion/api"
	"awslambdakinesisdeletion/awskinesis"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

)

//Handler function is used for deleting the kinesis stream, this function is triggered by an api request.
//function takes in request from api and returns a corresponding response and error.
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := api.ValidateAPIRequest(req) //validation of request
	if err != nil {
		return resp, nil
	}
	kinesisClient, resp, err := awskinesis.KinesisClientCreation() //kinesis client creation
	if err != nil {
		return api.Resp444, nil
	}
	streamName := req.PathParameters["streamName"]                         // stream name from path parameters from the url
	resp, err = awskinesis.DeleteKinesisStreams(kinesisClient, streamName) // process of stream deletion
	return resp, nil
}
func main() {

	lambda.Start(handler.Handler)
}
package handler_test

import (
	"awslambdakinesisdeletion/api"
	"awslambdakinesisdeletion/awssession"
	"awslambdakinesisdeletion/handler"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

//TestHandler is used for testing handler function
func TestHandlerWithInvalidPayload(t *testing.T) {
	awssession.ConfigDetails.Region = "us-west-2"
	req := events.APIGatewayProxyRequest{
		Resource: "jaswanth1",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod: "GET",
	}
	resp := api.Resp400
	response, _ := handler.Handler(req)
	if !reflect.DeepEqual(resp, response) {
		t.Error("test case for Handler function fails at TestHandler ", resp)
	}

}
func TestHandler(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		Resource: "jaswanth1",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod:     "GET",
		PathParameters: map[string]string{"streamName": "jaswanth"},
	}
	resp := api.Resp200
	response, _ := handler.Handler(req)
	if !reflect.DeepEqual(resp, response) {
		t.Error("test case for Handler function fails at TestHandler 2", response)
	}
}

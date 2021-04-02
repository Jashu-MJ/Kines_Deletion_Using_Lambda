package handler_test

import (
	"awslambdakinesiscreation/handler"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	req := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod: "POST",
		Body:       string(json.RawMessage(`{"name":"jaswanth","size":"1"}`)),
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode: 501,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: " Kinesis stream creation is not Successful",
	}
	response, _ := handler.Handler(req)
	assert.Equal(t, response, resp, "error occured")

}
func TestHandlerWithInvalidPayload(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod: "POST",
		Body:       string(json.RawMessage(`{"name":"jaswanth","size":""}`)),
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "Invalid payload",
	}
	response, _ := handler.Handler(req)
	assert.Equal(t, response, resp, "error occured")

}

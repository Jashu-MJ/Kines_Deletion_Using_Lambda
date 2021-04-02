package api_test

import (
	"awslambdakinesisdeletion/api"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestValidateAPIRequestWithGETMethod(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		PathParameters: map[string]string{
			"streamName": "Stream1",
		}}
	response := events.APIGatewayProxyResponse{}
	resp, _ := api.ValidateAPIRequest(request)
	assert.Equal(t, response, resp, "they should be equal")
}

func TestValidateAPIRequestWithPOSTMethod(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		PathParameters: map[string]string{
			"streamName": "Stream1",
		}}
	response := api.Resp400
	resp, _ := api.ValidateAPIRequest(request)
	assert.Equal(t, response, resp, "they should be equal")
}

func TestValidateAPIRequestWithGETMethodWithoutPathParams(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
	}
	response := api.Resp400
	resp, _ := api.ValidateAPIRequest(request)
	assert.Equal(t, response, resp, "they should be equal")
}

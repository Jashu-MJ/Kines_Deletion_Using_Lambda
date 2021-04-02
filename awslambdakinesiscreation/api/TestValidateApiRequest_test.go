package api_test

import (
	"awslambdakinesiscreation/api"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/magiconair/properties/assert"
)

func TestValidatingTheAPIRequestWithValidParams(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod: "POST",
		Body:       string(json.RawMessage(`{"name":"jaswanth","size":"1"}`)),
	}
	response := events.APIGatewayProxyResponse{StatusCode: 200, Body: "Unmarshall succesfull"}
	_, resp, _ := api.ValidatingTheAPIRequest(request)

	assert.Equal(t, response, resp, "error occured")

}

func TestValidatingTheAPIRequestWithSizeAsNil(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod: "POST",
		Body:       string(json.RawMessage(`{"name":"jaswanth","size":""}`)),
	}
	response := events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "Invalid payload",
	}
	_, resp, _ := api.ValidatingTheAPIRequest(request)

	assert.Equal(t, response, resp, "error occured")

}

func TestValidatingTheAPIRequestWithNameAsNull(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod: "POST",
		Body:       string(json.RawMessage(`{"name":"","size":"2",}`)),
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "Invalid payload",
	}
	_, resp, _ := api.ValidatingTheAPIRequest(request)

	assert.Equal(t, response, resp, "error occured")

}

func TestValidatingTheAPIRequestwithGETMethod(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod: "GET",
	}

	response := events.APIGatewayProxyResponse{StatusCode: 200, Body: "Invalid request HTTPmethod"}
	_, resp, _ := api.ValidatingTheAPIRequest(request)

	assert.Equal(t, response, resp, "error occured")

}

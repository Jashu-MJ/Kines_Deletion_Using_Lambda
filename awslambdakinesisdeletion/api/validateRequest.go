//Package api has validation function used for validating api requests
package api

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

// GET HTTP method
const getHTTPMethod string = "GET"

//ValidateAPIRequest is for validating whether the api request method is get
//method and ensures that the path parameters are not empty and return
//corresponding response and error
func ValidateAPIRequest(apiReq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if apiReq.HTTPMethod == getHTTPMethod && apiReq.PathParameters["streamName"] != "" { // HTTP Method is GET and path parameters are not null
		return events.APIGatewayProxyResponse{}, nil
	}
	return Resp400, errors.New("the htttp method is not GET method or params are empty")
}

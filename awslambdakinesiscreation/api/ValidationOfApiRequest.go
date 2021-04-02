// Package api has API Request Validation
package api

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// POST HTTP Method
const postHTTPMethod string = "POST"

//ValidatingTheAPIRequest takes in A request and process it according to the HTTP
//method and returns Api response and error if occurred
func ValidatingTheAPIRequest(apiReq events.APIGatewayProxyRequest) (RequestBody, events.APIGatewayProxyResponse, error) {
	var apiBody RequestBody
	var resp events.APIGatewayProxyResponse
	var err error
	if apiReq.HTTPMethod == postHTTPMethod {
		apiBody, resp, err = unmarshallingRequest(apiReq)
		if err != nil {
			log.Println(" request with valid HTTP method  eror occurred while unmarshall", err)
			return apiBody, resp, err
		} else if isReqBodyValid(apiBody) {
			log.Println("valid request and parse the body of request successfully ")
			return apiBody, resp, err

		} else {
			log.Println("valid request and parse the body of request successfully but insufficient params ")
			return apiBody, Resp400, errors.New("body of the json is not in the specified form or insufficient params")
		}
	}
	log.Println("Invalid request HTTPmethod")
	return apiBody, events.APIGatewayProxyResponse{StatusCode: 200, Body: "Invalid request HTTPmethod"}, errors.New("Invalid request HTTPmethod")
}

// Package api has unmarshlling the request function.
package api

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// UnmarshallingRequest functions takes in the request from the api gateway and decodes
// it in the form of ApiRequestBody if it is successful then returns error as nil or if
//not in  format then throws an error.Upon successful decoding the body it will be
//validating the body using the isReqBodyValid function which validates the body format
//and return a bool value and the function returns ApiRequestBody object, api response,
// error
func unmarshallingRequest(req events.APIGatewayProxyRequest) (RequestBody, events.APIGatewayProxyResponse, error) {
	var apiReq RequestBody
	err := json.Unmarshal([]byte(req.Body), &apiReq)
	if err != nil {
		log.Printf("Error occurred while decoding the Json or parsing the request body (%s); ...", err)
		return apiReq, Resp400, err
	}
	return apiReq, events.APIGatewayProxyResponse{StatusCode: 200, Body: "Unmarshall succesfull"}, err
}

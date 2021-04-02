//Package api has responses for corresponding occurrences
package api

import (
	"github.com/aws/aws-lambda-go/events"
)

// Resp200 is response with status code 200 when the kinesis stream creation is successful
var Resp200 events.APIGatewayProxyResponse = events.APIGatewayProxyResponse{
	StatusCode: 200,
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
	Body: "Created kinesis stream Succesfully",
}

// Resp501 is response with status code 501 when the kinesis stream creation is not successful
var Resp501 events.APIGatewayProxyResponse = events.APIGatewayProxyResponse{
	StatusCode: 501,
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
	Body: " Kinesis stream creation is not Successful",
}

// Resp400 is response with status code 400 when the payload in the request is invalid
var Resp400 events.APIGatewayProxyResponse = events.APIGatewayProxyResponse{
	StatusCode: 400,
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
	Body: "Invalid payload",
}

// Resp444 is response with status code 444 when the Aws session could not be established
var Resp444 events.APIGatewayProxyResponse = events.APIGatewayProxyResponse{
	StatusCode: 444,
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
	Body: "Aws session could not be established ",
}

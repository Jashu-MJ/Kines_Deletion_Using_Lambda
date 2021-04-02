package awsapigateway

import (
	"awslambdamonitor/awssession"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

//CreateAPIGatewayClient is used for creating aws api gateway client
//to communicate with api gateway.
func CreateAPIGatewayClient() (*apigateway.APIGateway, error) {
	awsSession, err := awssession.CreateSession()
	return apigateway.New(awsSession), err
}

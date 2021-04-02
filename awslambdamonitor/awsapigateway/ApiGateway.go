//Package awsapigateway has resource , method , endpoint integration and deployment
package awsapigateway

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
)

// ANY HTTP Method
const anyHTTPMethod = "ANY"

const authorizationNoneType = "NONE"

// POST HTTP Method
const postHTTPMethod = "POST"

// integaration type for the endpoint in api
const integrationAWSProxyType = "AWS_PROXY"

//createResource is Used for creating the resource with given resource name at
//given parent id and rest api id of REST API
func createResource(apiGatewayClient apigatewayiface.APIGatewayAPI, resourceName string, parentID string, restAPIID string) (*apigateway.Resource, error) {

	params := &apigateway.CreateResourceInput{ParentId: aws.String(parentID), PathPart: aws.String(resourceName), RestApiId: aws.String(restAPIID)}
	resource, err := apiGatewayClient.CreateResource(params)
	return resource, err
}

//putMethod used for creating the "ANY" method  under the given resource and
// rest api id
func putMethod(apiGatewayClient apigatewayiface.APIGatewayAPI, resourceID string, restAPIID string) (*apigateway.Method, error) {

	params := &apigateway.PutMethodInput{
		ApiKeyRequired:    aws.Bool(false),
		HttpMethod:        aws.String(anyHTTPMethod),
		ResourceId:        aws.String(resourceID),
		RestApiId:         aws.String(restAPIID),
		AuthorizationType: aws.String(authorizationNoneType),
	}
	method, err := apiGatewayClient.PutMethod(params)
	return method, err
}

//integrateWithLambda used for integrating the endpoint with the lambda function
func integrateWithLambda(apiGatewayClient apigatewayiface.APIGatewayAPI, rescID string, restAPIID string, lambdaFunctionName string) (*apigateway.Integration, error) {
	params := &apigateway.PutIntegrationInput{
		HttpMethod:            aws.String(anyHTTPMethod),
		ResourceId:            aws.String(rescID),
		RestApiId:             aws.String(restAPIID),
		Type:                  aws.String(integrationAWSProxyType),
		Uri:                   aws.String(lambdaFunctionName),
		IntegrationHttpMethod: aws.String(postHTTPMethod),
	}
	integration, err := apiGatewayClient.PutIntegration(params)
	return integration, err
}

//createDeployment used for deploying the rest api with stage name as "KinesisDeletion"
//at given rest api id.
func createDeployment(apiGatewayClient apigatewayiface.APIGatewayAPI, restAPIID string, stageName string) (*apigateway.Deployment, error) {
	params := &apigateway.CreateDeploymentInput{
		RestApiId: aws.String(restAPIID),
		StageName: aws.String(stageName),
	}
	deployment, err := apiGatewayClient.CreateDeployment(params)
	return deployment, err
}

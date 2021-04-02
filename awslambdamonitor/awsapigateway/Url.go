//Package awsapigateway has generation of link
package awsapigateway

import (
	"awslambdamonitor/awslambda"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

// stage name for deployment of api
const stageName = "KinesisDeletion"

// creating a path parameter for the resource with the stream name
const pathParameter = "{streamName}"

// arn of lambda function which is used for deletion of kinesis stream
const arnOfDeleteFunction = "arn:aws:lambda:us-west-2:225327301834:function:chartsearch-HiEng-JM078255-kinesis-deletion-test"

//CreationOfURL is used for creating a url in the name of the kinesis stream and this url can be used for deleting the kinesis stream
func CreationOfURL(apiGatewayClient apigatewayiface.APIGatewayAPI, lambdaClient lambdaiface.LambdaAPI) (string, error) { // *apigateway.APIGateway

	resourceName := os.Getenv("StreamName")      //creating a resource with stream name as resource name
	parentID := os.Getenv("ParentID")            //parent id is the id of the parent resource in API
	restAPIID := os.Getenv("RESTAPI_ID")         //id of REST API
	lambdaFunctionName := os.Getenv("LambdaARN") // lambda function at end point of the api which has to be triggered.
	var link string = "https://" + restAPIID + ".execute-api." + os.Getenv("AwsRegion") + ".amazonaws.com/" + stageName + "/" + os.Getenv("StreamName") + "/" + os.Getenv("StreamName")

	isTrue, err := isURLAlreadyExists(apiGatewayClient, resourceName, restAPIID)
	if err != nil {
		return "", err
	}
	if isTrue {
		return link, err
	}

	// creating the resource with stream name
	resource, err := createResource(apiGatewayClient, resourceName, parentID, restAPIID)
	if err != nil {
		return "", err
	}

	//creating the path parameter under the resource with resource name as stream name
	resourceOfPathParam, err := createResource(apiGatewayClient, pathParameter, *resource.Id, restAPIID)
	if err != nil {
		return "", err
	}

	//putting ANY  HTTP Method under the path parameter resource.
	_, err = putMethod(apiGatewayClient, *resourceOfPathParam.Id, restAPIID)
	if err != nil {
		return "", err
	}

	// integrating the endpoint of the resource with the lambda function.
	_, err = integrateWithLambda(apiGatewayClient, *resourceOfPathParam.Id, restAPIID, lambdaFunctionName)
	if err != nil {
		return "", err
	}

	// granting permissions for the api to invoke the lambda function at a given ARN
	err = awslambda.GivePermissionsForIntegration(lambdaClient, os.Getenv("StreamName"), arnOfDeleteFunction)
	if err != nil {

		return "", err
	}

	// deploying the rest api so that it will be functional .
	_, err = createDeployment(apiGatewayClient, restAPIID, stageName)

	return link, err

}
func isURLAlreadyExists(apiGatewayClient apigatewayiface.APIGatewayAPI, pathPart string, restAPIID string) (bool, error) {

	params := &apigateway.GetResourcesInput{
		Limit:     aws.Int64(500),
		RestApiId: aws.String(restAPIID),
	}

	resrcs, err := apiGatewayClient.GetResources(params)
	if err != nil {
		return false, err
	}
	for _, resrc := range resrcs.Items {
		resource := *resrc
		if resource.PathPart == nil { // if resource is nill then continues with next resource
			continue
		}
		if pathPart == *resource.PathPart {
			return true, err
		}
	}
	return false, err
}
